import { describe, expect, it } from "vitest";

import { deepToCamelCase } from ".";

describe("utils", () => {
	it("deepToCamelCase", () => {
		expect(
			deepToCamelCase({
				a_b: {
					c_d: {
						a: 1,
						b: 2,
					},
				},
			}),
		).toMatchInlineSnapshot(`
			{
			  "aB": {
			    "cD": {
			      "a": 1,
			      "b": 2,
			    },
			  },
			}
		`);
	});
});
