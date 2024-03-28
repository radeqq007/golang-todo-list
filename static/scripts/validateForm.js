function validate() {
  const username = document.getElementById('username').value;
  const password = document.getElementById('password').value;
  const error = document.getElementById('error');

  error.innerHTML = '';

  if (username.length == 0 || password.length == 0) {
    error.innerHTML = 'Please fill in all fields';
    return false;
  }

  const regex = /^[a-zA-Z0-9_]+$/;
  if (!regex.test(username)) {
    error.innerHTML =
      'Username can only contain letters, numbers and underscores';
    return false;
  }

  if (password.length < 8) {
    error.innerHTML = 'Password must be at least 8 characters long';
    return false;
  }

  return true;
}
