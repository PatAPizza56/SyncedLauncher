function btnOutline(parentId, text, textClass, onClick) {
  var btn = document.createElement("div");

  btn.className = "btn outline";
  btn.innerHTML = `<span class="${textClass}">${text}</span>`;
  btn.addEventListener("click", onClick);

  document.getElementById(parentId).appendChild(btn);
}

function btnFill(parentId, text, textClass, onClick) {
  var btn = document.createElement("div");

  btn.className = "btn fill";
  btn.innerHTML = `<span class="${textClass}">${text}</span>`;
  btn.addEventListener("click", onClick);

  document.getElementById(parentId).appendChild(btn);
}

function btnOutlineLink(parentId, text, textClass, href) {
  var btn = document.createElement("a");

  btn.className = "btn outline";
  btn.innerHTML = `<span class="${textClass}">${text}</span>`;
  btn.href = href;

  document.getElementById(parentId).appendChild(btn);
}

function btnFillLink(parentId, text, textClass, href) {
  var btn = document.createElement("a");

  btn.className = "btn fill";
  btn.innerHTML = `<span class="${textClass}">${text}</span>`;
  btn.href = href;

  document.getElementById(parentId).appendChild(btn);
}
