import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path";


// https://vitejs.dev/config/
export default defineConfig({
  server: {
    // proxy: {
    //   '/user': {
    //     target: 'http://localhost',
    //     changeOrigin: true,
    //     secure: false,
    //     configure: (proxy, options) => {
    //       proxy.on('proxyRes', (proxyRes, req, res) => {
    //         proxyRes.headers['Access-Control-Allow-Origin'] = 'http://localhost:5555';
    //         proxyRes.headers['Access-Control-Allow-Methods'] = 'GET, POST, PUT, DELETE, OPTIONS';
    //         proxyRes.headers['Access-Control-Allow-Headers'] = 'Content-Type, Authorization';
    //         proxyRes.headers['Access-Control-Allow-Credentials'] = 'true';
    //       });
    //     }
    //   }
    // },
    // 后端CORS
    port: 5555,
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
      "components": path.resolve(__dirname, "src/components"),
      "styles": path.resolve(__dirname, "src/styles"),
      "plugins": path.resolve(__dirname, "src/plugins"),
      "views": path.resolve(__dirname, "src/views"),
      "layouts": path.resolve(__dirname, "src/layouts"),
      "utils": path.resolve(__dirname, "src/utils"),
      "apis": path.resolve(__dirname, "src/apis"),
      "dirs": path.resolve(__dirname, "src/directives"),
    },
  },
  plugins: [vue()],
})