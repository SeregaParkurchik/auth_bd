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


curl -X 'POST' \
  'http://localhost:8080/house/create' \
  -H 'accept: application/json' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7ImVtYWlsIjoic2VyZWdhQG1haWwucnUiLCJ1c2VyX3R5cGUiOiJjbGllbnQifSwiaWF0IjoxNzM4NzQzMzQzLCJleHAiOjE3Mzg3ODY1NDN9.xzOf2LrX7Gh4JAYczY3dorryuc3MHF6585YSrXjLtfo' \
  -H 'Content-Type: application/json' \
  -d '{
  "address": "Лесная улица, 7, Москва, 125196",
  "year": 2000,
  "developer": "Мэрия города"
}'