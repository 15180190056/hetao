function login() {
    const username = document.getElementById('username').value;
    const password = document.getElementById('password').value;
    const messageElement = document.getElementById('message');
    const form = new FormData();

    form.append('username', username);
    form.append('password', password);

    // 发送POST请求到后端地址
    fetch('http://127.0.0.1:8083/v1/login', {
        method: 'POST',
        body: form
    })
        .then(response => {
            if (!response.ok) {
                throw new Error('Network response was not ok');
            }
            return response.json();
        })
        .then(data => {
            if (data.status === 'success') {
                messageElement.innerHTML = data.message;
                messageElement.style.backgroundColor = 'lightgreen';
                // 登录成功后可以进行页面跳转或其他操作
            } else {
                messageElement.innerHTML = 'Invalid username or password';
                messageElement.style.backgroundColor = 'lightcoral';
            }
        })
        .catch(error => {
            console.error('There has been a problem with your fetch operation:', error);
            messageElement.innerHTML = 'An error occurred during login';
            messageElement.style.backgroundColor = 'lightcoral';
        });
}