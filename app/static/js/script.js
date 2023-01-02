window.onload = function () {
    const div1 = document.getElementById("block1");
    const div2 = document.getElementById("block2");
    div1.addEventListener('mouseover', dragElement(div1));
    div2.addEventListener('mouseover', dragElement(div2));

    function dragElement(elmnt) {

        var pos1 = 0, pos2 = 0, pos3 = 0, pos4 = 0;

        if (document.getElementById(elmnt.id + "h")) {
          // if present, the header is where you move the DIV from:
          document.getElementById(elmnt.id + "h").onmousedown = dragMouseDown;
        } else {
          // otherwise, move the DIV from anywhere inside the DIV:
          elmnt.onmousedown = dragMouseDown;
        }

        function dragMouseDown(e) {
            e = e || window.event;
            e.preventDefault();
            // get the mouse cursor position at startup:
            pos3 = e.clientX;
            pos4 = e.clientY;
            setDim(document.getElementById("block1"), document.getElementById("block2"));
            document.getElementById("block1").style.position = "absolute";
            document.getElementById("block2").style.position = "absolute";
            document.onmouseup = closeDragElement;
            // call a function whenever the cursor moves:
            document.onmousemove = elementDrag;
        }

        function elementDrag(e) {
            e = e || window.event;
            e.preventDefault();
            // calculate the new cursor position:
            pos1 = pos3 - e.clientX;
            pos2 = pos4 - e.clientY;
            pos3 = e.clientX;
            pos4 = e.clientY;
            // set the element's new position:
            const rect = elmnt.getBoundingClientRect();
            elmnt.style.left = (rect.left - pos1 + window.scrollX) + "px";
            elmnt.style.top = (rect.top - pos2 + window.scrollY) + "px";
        }

        function closeDragElement() {
            // stop moving when mouse button is released:
            document.onmouseup = null;
            document.onmousemove = null;
        }
        function setDim(el1, el2) {
            const rect1 = el1.getBoundingClientRect(), rect2 = el2.getBoundingClientRect();
            el1.style.left = (rect1.left + window.scrollX) + "px";
            el1.style.top = (rect1.top + window.scrollY) + "px";
            el2.style.left = (rect2.left + window.scrollX) + "px";
            el2.style.top = (rect2.top + window.scrollY) + "px";
            el1.style.width = (rect1.width-10) + "px";
            el1.style.height = (rect1.height-10) + "px";
            el2.style.width = (rect2.width-10) + "px";
            el2.style.height = (rect2.height-10) + "px";
            console.log(rect1, rect2);
        }
    }
};

