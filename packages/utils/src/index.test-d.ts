import { describe, expectTypeOf, it } from "vitest";

import { deepToCamelCase } from ".";

describe("utils", () => {
	it("deepToCamelCase", () => {
		expectTypeOf(
			deepToCamelCase({
				a_b: {
					c_d: {
						a: 1,
						b: 2,
					},
				},
			}),
		).toMatchTypeOf<{
			aB: {
				cD: {
					a: number;
					b: number;
				};
			};
		}>();
	});
});
