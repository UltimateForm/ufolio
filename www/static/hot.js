const source = new EventSource("/hot");
source.addEventListener("open", () => {
  console.log("Connected to hot reload server");
});

source.addEventListener("error", (e) => {
  console.error("Hot reload server error:", e);
});

source.addEventListener("message", (e) => {
  console.log("Hot reload message:", e.data);
});
