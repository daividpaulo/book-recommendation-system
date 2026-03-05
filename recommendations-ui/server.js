import express from "express";

const app = express();
const port = process.env.PORT || 3000;
const apiBase = process.env.API_URL || "http://recommendations-api:8080";

app.use(express.static("public"));
app.use(express.json());

app.get("/health", (_, res) => {
  res.json({ status: "ok", service: "recommendations-ui" });
});

app.post("/proxy/*", async (req, res) => {
  const path = req.params[0];
  const response = await fetch(`${apiBase}/${path}`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(req.body || {})
  });
  const body = await response.text();
  res.status(response.status).send(body);
});

app.get("/proxy/*", async (req, res) => {
  const path = req.params[0];
  const response = await fetch(`${apiBase}/${path}`);
  const body = await response.text();
  res.status(response.status).send(body);
});

app.listen(port, () => {
  console.log(`recommendations-ui listening on :${port}`);
});
