

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

function setLowFreq() {
    el = getData();

    if (el.Freq > 1 ) {
        el.Freq -= 1
    }

    let freqq = el.Freq;
    let idd = el.Id;

    let response = fetch('http://localhost:5500/set', { method: 'post', body: JSON.stringify({ freqq, idd }) })
}

function setHiFreq() {
    el = getData();

    if (el.Freq < 5) {
        el.Freq += 1
    }

    let freqq = el.Freq;
    let idd = el.Id;

    let response = fetch('http://localhost:5500/set', { method: 'post', body: JSON.stringify({ freqq, idd }) })
}

function getData() {
    let response = fetch('http://localhost:5500/get', { method: 'post', headers: { "Content-Type": "application/json" } })
        .then((resp) => resp.json())
        .then(function (data) {
            elm = data;
            /* randDataEl = data[Math.floor(Math.random() * data.length)] */
            document.getElementById('one').innerHTML = data.Fword;
            document.getElementById('answer').innerHTML = data.Sword;
        })
        .catch(function (error) {
            console.log(error)
        })

    /* document.getElementById('one').innerHTML = randDataEl.Fword;
    document.getElementById('answer').innerHTML = randDataEl.Sword; */

    s = document.getElementById("twobtn").style;
    if (s.backgroundColor != 'black') {
        s.backgroundColor = 'black';
    }

    return elm
}

function HiddenLearn(Element) {
    if (Element.style.opacity == 1) {
        Element.style.opacity = 0;
    } else {
        Element.style.opacity = 1;
    }
}

document.getElementById("get").addEventListener("click", getData);
document.getElementById("get1").addEventListener("click", setLowFreq);
document.getElementById("get2").addEventListener("click", setHiFreq);