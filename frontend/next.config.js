const withWorkbox = require("next-with-workbox");

module.exports = withWorkbox({
    workbox: {
        swSrc: "worker.js",
        force: true
    }
})