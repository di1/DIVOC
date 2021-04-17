const certificateController = require("./src/routes/certificate_controller");
const {initKafa} = require("./src/services/kafka_service");

const http = require('http');
const {KeycloakFactory} = require("./src/services/keycloak_service");

const port = process.env.PORT || 4321;

const server = http.createServer(async (req, res) => {
    console.time(req.url)
    console.log(`API ${req.method} ${req.url} called`);
    if (req.method === 'GET' && req.url.startsWith("/certificate/api/certificate/")) {
        const data = await certificateController.getCertificate(req, res);
        res.end(data)
    } else {
        res.end(`{"error": "${http.STATUS_CODES[404]}"}`)
    }
    console.timeEnd(req.url)
});

server.listen(port, async () => {
    await initKafa();
    console.log(`Server listening on port ${port}`);
});