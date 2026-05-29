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
    if (elmnt.getAttribute("aria-expanded") === "true") return;
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

// magic
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
      document.querySelectorAll(".window[aria-current]").forEach((w) => {
        w.removeAttribute("aria-current");
      });
      win.setAttribute("aria-current", "page");
    });
    win.addEventListener("mousedown", function () {
      if (!win.hasAttribute("aria-current")) {
        win.focus();
      }
    });

    win
      .querySelectorAll('[aria-label="Close"], [aria-label="Minimize"]')
      .forEach((el) => {
        const controller = document.querySelector(
          `[aria-controls="${win.id}"]`,
        );
        el.addEventListener("click", function () {
          win.blur();
          win.setAttribute("aria-hidden", true);
          if (controller) {
            controller.setAttribute("aria-pressed", false);
          }
        });
      });

    const maximizeBtn = win.querySelector('[aria-label="Maximize"]');
    if (!maximizeBtn) return;
    function maximize() {
      const isExpanded = win.getAttribute("aria-expanded") === "true";
      win.setAttribute("aria-expanded", !isExpanded);
    }
    win.querySelector(".title-bar")?.addEventListener("dblclick", function () {
      maximize();
    });
    maximizeBtn.addEventListener("click", function () {
      maximize();
    });
  });

  document.querySelectorAll("button[aria-controls]").forEach((btn) => {
    btn.addEventListener("click", function () {
      const panelId = btn.getAttribute("aria-controls");
      const panel = document.getElementById(panelId);
      const isRadio = btn.getAttribute("role") === "radio";
      const btnStateKey = isRadio ? "aria-checked" : "aria-pressed";
      const isOpen = btn.getAttribute(btnStateKey) === "true";
      btn.setAttribute(btnStateKey, !isOpen);

      if (panel) {
        panel.setAttribute("aria-hidden", isOpen);
        if (!isOpen && panel.classList.contains("window")) {
          panel.focus();
          // auto-maximize on mobile
          if (window.innerWidth < 768) {
            panel.setAttribute("aria-expanded", "true");
          }
        }
      }

      if (isRadio) {
        const closestRadioGroup = btn.closest("[role=radiogroup]");
        if (!closestRadioGroup) {
          console.error("radio button is not inside a radio group", btn);
          return;
        }
        closestRadioGroup
          .querySelectorAll(
            `button[role=radio][aria-checked=true]:not([aria-controls=${panelId}])`,
          )
          .forEach((radioBtn) => {
            radioBtn.setAttribute(btnStateKey, false);
            const disableTarget = radioBtn.getAttribute("aria-controls");
            if (disableTarget) {
              document
                .getElementById(disableTarget)
                ?.setAttribute("aria-hidden", true);
            }
          });
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

  const skillsTree = document.querySelector("#skills-tree");
  const treeToggleBtn = document.querySelector("#tree-toggle");
  // this is just for mobile version of about-tech window, runs in desktop but low cost no-op
  skillsTree.querySelectorAll("button[aria-controls]").forEach((btn) => {
    btn.addEventListener("click", () => {
      if (treeToggleBtn && window.innerWidth < 768) {
        treeToggleBtn.click();
      }
    });
  });
});
