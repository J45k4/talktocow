import react from "@vitejs/plugin-react"
import { defineConfig } from "vitest/config"

export default defineConfig({
	plugins: [react()],
	test: {
		environment: "jsdom",
		globals: true,
	},
	server: {
		port: 3080,
		proxy: {
			"/api": "http://127.0.0.1:12001",
			"/ws": {
				target: "ws://127.0.0.1:12001",
				ws: true,
			},
		},
	},
	preview: {
		port: 3080,
	},
})
