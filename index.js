function toggleMenu() {
  console.log("toggling menu...");
  var elem = document.getElementById("menu-items");
  if (elem.classList.contains("open-menu")) {
    elem.classList.remove("open-menu");
  } else {
    elem.classList.add("open-menu");
  }
}
