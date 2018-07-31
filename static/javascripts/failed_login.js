window.onload = () => {
  // Set the onClock event of the button.
  $("form").on("submit", () => {
    event.preventDefault();
    console.log("Ya got me.");
    let username = $("#username").val();
    let password = $("#password").val();
    let url = "/documents?username="+ username + "&password="+ password;
    console.log(username);
    console.log(password);
    console.log(url);
    location.replace(url);
  });
};
