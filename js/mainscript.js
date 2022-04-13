let xttp = new XMLHttpRequest();



function ShowTwoBtnText(Element) {
    Element.style.backgroundColor = 'white';
}

function HiddenTwoBtnText(n) {
    s = document.getElementById(n).style;
    s.backgroundColor = 'black';

    xttp.addEventListener("load", "Next()")
    xttp.open("get", "main.go", true);
    xttp.send();
}

function HiddenLearn(Element) {
    if (Element.style.opacity == 1) {
        Element.style.opacity = 0;
    } else {
        Element.style.opacity = 1;
    }
}