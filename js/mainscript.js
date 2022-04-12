function ShowTwoBtnText(Element) {
    Element.style.backgroundColor = 'white';
}

function HiddenTwoBtnText(n) {
    s = document.getElementById(n).style;
    s.backgroundColor = 'black';
}

function HiddenLearn(Element) {
    if (Element.style.opacity == 1) {
        Element.style.opacity = 0;
    } else {
        Element.style.opacity = 1;
    }
}