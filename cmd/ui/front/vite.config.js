module.exports = {
    outDir: "../dist",
    proxy: {
        "/api/data": "http://localhost:8080/api/data"
    }
}