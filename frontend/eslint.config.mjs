import js from "@eslint/js"
import globals from "globals"
import tseslint from "typescript-eslint"

export default tseslint.config(
	js.configs.recommended,
	...tseslint.configs.recommended,
	{
		files: ["**/*.{ts,tsx}"],
		languageOptions: {
			ecmaVersion: "latest",
			globals: {
				...globals.browser,
				...globals.node,
			},
		},
		rules: {
			"@typescript-eslint/no-explicit-any": "off",
			"@typescript-eslint/no-unused-vars": "off",
			"no-unused-vars": "off",
			"no-empty": "off",
			"no-mixed-spaces-and-tabs": "off",
			"no-var": "off",
			"prefer-const": "off",
		},
	},
	{
		ignores: ["dist", "node_modules"],
	}
)
