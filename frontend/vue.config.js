// Option 2: Proxy tutto il traffico con percorso iniziale "/api" a http://localhost:8080
module.exports = {
  devServer: {
    proxy: {
      "^/api": {
        target: "http://localhost:8081",
        changeOrigin: true,
      },
    },
  },
};
