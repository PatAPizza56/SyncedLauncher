function nav(parentId) {
  var nav = document.createElement("div");
  nav.innerHTML = `
  <div class="nav">
    <svg class="nav-logo" viewBox="0 0 46 26">
      <path
        class="nav-logo-path"
        d="M23 13L13 23L3 13L13 3L23 13ZM23 13L33 3L43 13L33 23L23 13Z"
      />
      <defs>
        <linearGradient id="linear" x1="153.823" y1="4.92392" x2="153.823" y2="383.5" gradientUnits="userSpaceOnUse">
          <stop stop-color="#0082FB"/>
          <stop offset="1" stop-color="#0019FB"/>
        </linearGradient>
      </defs>
    </svg>
    <div class="nav-buttons">
      <a class="nav-button" href="/home">
        <span class="p">Home</span>
      </a>
      <a class="nav-button" href="/shop">
        <span class="p">Shop</span>
      </a>
      <a class="nav-button" href="/games">
        <span class="p">Games</span>
      </a>
    </div>
    <a class="nav-cta" href="/register">
      <span class="small">Sign up</span>
    </a>
  </div>
  `;

  document.getElementById(parentId).appendChild(nav);
}
