// Get the modal
var modal = document.getElementById("modal");
var modal__overlay = document.getElementById("modal__overlay");
// Get the register that opens the modal
var btnRegister = document.querySelectorAll('.register');
var btnLogin = document.querySelectorAll('.login');
var btnBack = document.querySelectorAll('.back');

window.onclick = function(event) {
    if (event.target == modal__overlay) {
      modal.style.display = "none";
    }
  }

btnRegister.forEach(element => {
    element.addEventListener("click", ()=>{
        modal.style.display = "flex";
        login_form.style.display = "none";
        register_form.style.display = "block";
    });
});

btnLogin.forEach(element => {
    element.addEventListener("click", ()=>{
        modal.style.display = "flex";
        login_form.style.display = "block";
        register_form.style.display = "none";
    });
});

btnBack.forEach(element => {
    element.addEventListener("click", ()=>{
        modal.style.display = "none";
    });
});
