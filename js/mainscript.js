

function ShowTwoBtnText(Element) {
    Element.style.backgroundColor = 'white';
}




function HiddenTwoBtnText(n) {
    console.log(123)
    let response = fetch('http://localhost:5500/get');
    console.log(response)

    if (response.ok) {
        let data = response.json();
        console.log(data)

        // document.getElementById(n).innerHTML =      
    } else {
        alert("cant get word")
    }

    s = document.getElementById(n).style;
    s.backgroundColor = 'black';
}




function getData() {
    let response = fetch('http://localhost:5500/get', { method: 'post', headers: { "Content-Type": "application/json" }, body: JSON.stringify({ "test": "test" }) })
        .then((resp) => resp.json())
        .then(function (data) {
            document.getElementById('one').innerHTML = data.Fword;
            document.getElementById('answer').innerHTML = data.Sword;
        })
        .catch(function (error) {
            console.log(error)
        })

    s = document.getElementById("twobtn").style;
    if (s.backgroundColor != 'black') {
        s.backgroundColor = 'black';
    }
}


function HiddenLearn(Element) {
    if (Element.style.opacity == 1) {
        Element.style.opacity = 0;
    } else {
        Element.style.opacity = 1;
    }
}

document.getElementById("get").addEventListener("click", getData);
document.getElementById("get1").addEventListener("click", getData);
document.getElementById("get2").addEventListener("click", getData);