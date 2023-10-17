const $Proyect = document.querySelector("$myProyect");

const changeProyect = () =>{
    console.log("Change");
}
const agregar = ()=>{
}

$Proyect.addEventListener("change", changeProyect)

const show = () => {
    const id = $Proyect.selectedIndex;
    if(id==-1)return;
    const selectedopcion = $Proyect.opcions[id];
    
}

