function input(parentId, heading, warning, description) {
  var input = document.createElement("div");
  input.innerHTML = `
  <div class="input">
    <span class="input-label">
      <p class="input-label-heading">
        <b>${heading}</b>
      </p>
      <small id="${parentId}Warning" class="input-label-warning">${warning}</small>
    </span>
    <input id="${parentId}Field" class="p input-text" type="text" />
    <small class="input-description">${description}</small>
  </div>
  `;

  document.getElementById(parentId).appendChild(input);
}
