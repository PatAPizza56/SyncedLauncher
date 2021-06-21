function CounterComponent() {
  let state = {
    count: 0,
  };

  let incrementCounter = () => {
    state.count += 1;
  };

  let updateButton = (button) => {
    button.innerText = state.count;
  };

  let renderButton = () => {
    let button = document.createElement("button");
    button.addEventListener("click", () => {
      incrementCounter();
      updateButton(button);
    });

    updateButton(button);
    return button;
  };

  let renderContainer = () => {
    let div = document.createElement("div");
    div.appendChild(renderButton());

    return div;
  };

  return renderContainer();
}

//document.body.appendChild(CounterComponent());
