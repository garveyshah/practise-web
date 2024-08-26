const buttons = document.querySelectorAll("button");

for (const button of buttons) {
button.addEventListener("click", updateName);
}


for (const button of buttons) {
    button.addEventListener("click",createParagraph);
}
function updateName() {
    const name = prompt("Enter a new name");
    button.textContent = `player 1: ${name}`;
}

function createParagraph() {
    const para = document.createElement("p");
    para.textContent = "You clicked the button!"
    document.body.appendChild(para);
}
