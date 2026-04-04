// snatched from https://www.w3schools.com/howto/howto_js_draggable.asp and tweaked a bit
function dragElement(elmnt) {
  var pos1 = 0,
    pos2 = 0,
    pos3 = 0,
    pos4 = 0;

  const titleBar = elmnt.querySelector(".title-bar");
  if (!titleBar) {
    return;
  }
  titleBar.onmousedown = dragMouseDown;

  function dragMouseDown(e) {
    e = e || window.event;

    if (e.target !== titleBar) {
      // why? we only want to drag the window when the title bar is clicked, not any other element inside the window
      // so we return early if the clicked element is not the title bar
      // the other approach would to do something like `if (e.target.closest('.title-bar-controls')) return;`
      // but i find it cleaner to check if the target is the title bar directly
      return;
    }

    e.preventDefault();
    // get the mouse cursor position at startup:
    pos3 = e.clientX;
    pos4 = e.clientY;
    document.onmouseup = closeDragElement;
    // call a function whenever the cursor moves:
    document.onmousemove = elementDrag;

    // titleBarText.textContent = `X: ${pos3}, Y: ${pos4}`;
  }

  function elementDrag(e) {
    e = e || window.event;
    e.preventDefault();
    // calculate the new cursor position:
    pos1 = pos3 - e.clientX;
    pos2 = pos4 - e.clientY;
    pos3 = e.clientX;
    pos4 = e.clientY;
    // set the element's new position:
    elmnt.style.top = elmnt.offsetTop - pos2 + "px";
    elmnt.style.left = elmnt.offsetLeft - pos1 + "px";
    // titleBarText.textContent = `X: ${pos1}, Y: ${pos2}`;
  }

  function closeDragElement() {
    // stop moving when mouse button is released:
    document.onmouseup = null;
    document.onmousemove = null;
  }
}

function textToRgb(repoName) {
  let hash = 0;
  for (let i = 0; i < repoName.length; i++) {
    hash = (hash << 5) - hash + repoName.charCodeAt(i);
    hash = hash & hash;
  }

  const r = (hash >> 0) & 255;
  const g = (hash >> 8) & 255;
  const b = (hash >> 16) & 255;

  return {
    bg: `rgb(${r}, ${g}, ${b})`,
    text: `rgb(${255 - r}, ${255 - g}, ${255 - b})`,
  };
}
document.addEventListener("DOMContentLoaded", function () {
  const wins = document.querySelectorAll(".window:not(.start-menu)");

  wins.forEach(function (win) {
    dragElement(win);
    win.addEventListener("focus", function () {
      win.parentElement.append(win);
    });
    win.addEventListener("mousedown", function () {
      win.focus();
    });
  });

  document.querySelectorAll("button[aria-controls]").forEach((btn) => {
    btn.addEventListener("click", function () {
      const panelId = btn.getAttribute("aria-controls");
      const panel = document.getElementById(panelId);
      const isOpen = btn.getAttribute("aria-expanded") === "true";
      btn.setAttribute("aria-expanded", !isOpen);
      panel.setAttribute("aria-hidden", isOpen);
      if (!isOpen) {
        panel.focus();
      }
    });
  });

  document.querySelectorAll(".repo").forEach((repo) => {
    const name = repo.getAttribute("data-repo-name");
    const svg = repo.querySelector("svg");
    if (svg) {
      const colors = textToRgb(name);
      svg.style.setProperty("--bg-color", colors.bg);
      svg.style.setProperty("--icon-color", colors.text);
    }
  });
});
