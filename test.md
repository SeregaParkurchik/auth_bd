fetch('http://localhost:8080/register', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify({
        email: 'maria@mail.ru',
        password: '2000',
        user_type: 'client'
    })
})
.then(response => response.json())
.then(data => console.log(data))
.catch(error => console.error('Ошибка:', error));

=============

fetch('http://localhost:8080/api/login', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json'
    },
    body: JSON.stringify({
        email: 'serega@mail.ru', // Здесь вы можете использовать email или user ID
        password: '123456'
    })
})
.then(response => {
    if (!response.ok) {
        throw new Error('Ошибка при аутентификации');
    }
    return response.json();
})
.then(data => console.log(data)) // Здесь вы получите токен
.catch(error => console.error('Ошибка:', error));