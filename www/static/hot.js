document.addEventListener("DOMContentLoaded", () => {
  const source = new EventSource("/hot/live");
  console.log("hot reload EventSource created");

  source.addEventListener("open", () => {
    // something is up here
    // for whatever reason this never getting called, figure out later
    console.log("connected to hot reload server");
  });

  source.addEventListener("error", (e, o) => {
    console.error("hot reload event source error:", e, o);
  });

  source.addEventListener("message", (e) => {
    if (e.data === "RELOAD") {
      location.reload();
    }
  });
});
