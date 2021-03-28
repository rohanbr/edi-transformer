package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/jf-tech/omniparser"
	"github.com/jf-tech/omniparser/transformctx"
)

func main() {
	input := `ISA*00*          *00*          *02*CPC            *ZZ*00602679321    *191103*1800*U*00401*000001644*0*P*>
	GS*QM*CPC*00602679321*20191103*1800*000001644*X*004010
	ST*214*000000001
	B10*4343638097845589              *4343638097845589              *CPCC
	L11*4343638097845589*97
	L11*0000*86`

	schema := `{
		"parser_settings": {
			"version": "omni.2.1",
			"file_format_type": "edi"
		},
		"file_declaration": {
			"segment_delimiter": "|",
			"element_delimiter": "*",
			"component_delimiter": ">",
			"ignore_crlf": true,
			"segment_declarations": [
			]
		},
		"segment_declarations": [
		{
			"name": "ISA",
			"child_segments": [
				{
					"name": "GS",
					"child_segments": [
						{
							"name": "invoiceInfo", "type": "segment_group", "is_target": true,
							"child_segments": [
								{ "name": "ST" },
								{ "name": "B3" },
								{ "name": "C3" },
								{ "name": "ITD" },
								{ "name": "N9" },
								{
									"name": "partyInfo", "type": "segment_group",
									"child_segments": [
										{ "name": "N1" },
										{ "name": "N2" },
										{ "name": "N3" },
										{ "name": "N4" },
										{ "name": "N9" }
									]
								},
								{
									"name": "lineItemInfo", "type": "segment_group",
									"child_segments": [
										{ "name": "LX" },
										{ "name": "N9" },
										{ "name": "L5" },
										{ "name": "L0" },
										{ "name": "L1" },
										{ "name": "L4" },
										{
											"name": "consigneeInfo", "type": "segment_group",
											"child_segments": [
												{ "name": "N1" },
												{ "name": "N2" },
												{ "name": "N3" },
												{ "name": "N4" },
												{ "name": "N9" },
												{
													"name": "cartonInfo", "type": "segment_group",
													"child_segments": [
														{ "name": "CD3" },
														{ "name": "N9" }
													]
												}
											]
										}
									]
								},
								{ "name": "L3" },
								{ "name": "SE" }
							]
						}
					]
				},
				{ "name": "GE" }
			]
		},
		{ "name": "IEA" }
	],
		"transform_declarations": {
			"FINAL_OUTPUT": { "object": {
				"invoice_number": { "xpath": "B3/invoiceNumber" }
			}}
		}
	}`

	schem, err := omniparser.NewSchema("your schema name", strings.NewReader(schema))

	if err != nil {
		fmt.Println("this is dsfs")
		fmt.Println(err.Error())
		return
	}
	transform, err := schem.NewTransform("your input name", strings.NewReader(input), &transformctx.Ctx{})
	for {
		output, err := transform.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
		}
		// output contains a []byte of the ingested and transformed record.
		fmt.Println(string(output))
		// Also transform.RawRecord() gives you access to the raw record.
		a, err := transform.RawRecord()
		fmt.Println(a.Checksum())
	}
}
