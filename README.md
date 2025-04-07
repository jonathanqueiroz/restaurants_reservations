# Restaurants Reservations

Este é um projeto de gerenciamento de reservas de restaurantes, desenvolvido com uma arquitetura de microsserviços utilizando **Clean Architecture** e **RabbitMQ** para comunicação assíncrona entre os serviços.

---

## 🚀 Tecnologias Utilizadas

- **Linguagem**: [Node.js](https://nodejs.org/) (ou substitua pela linguagem usada)
- **Framework**: [Express.js](https://expressjs.com/) (ou outro framework utilizado)
- **Mensageria**: [RabbitMQ](https://www.rabbitmq.com/)
- **Banco de Dados**: [PostgreSQL](https://www.postgresql.org/) (ou outro banco utilizado)
- **Arquitetura**: Clean Architecture
- **Padrões Adotados**:
  - **DDD (Domain-Driven Design)**
  - **SOLID**
  - **Event-Driven Architecture**
  - **RESTful APIs**

---

## 📚 Estrutura do Projeto

O projeto segue os princípios da **Clean Architecture**, com as seguintes camadas principais:

- **Domain**: Contém as regras de negócio e entidades.
- **Application**: Contém os casos de uso e lógica de aplicação.
- **Infrastructure**: Contém a implementação de serviços externos, como banco de dados e mensageria.
- **Interface (API)**: Contém os controladores e rotas para interação com o sistema.

---

## 📖 Endpoints e Rotas

### **1. Reservas**
#### **POST /reservations**
- **Descrição**: Cria uma nova reserva.
- **Body**:
  ```json
  {
    "restaurant_id": 1,
    "user_id": 1,
    "date": "2025-04-10",
    "time": "20:00:00"
  }
  ```

#### **GET /reservations/:id**
- **Descrição**: Retorna os detalhes de uma reserva específica.
- **Resposta**:
  ```json
  {
    "id": 1,
    "restaurant_id": 1,
    "user_id": 1,
    "date": "2025-04-10",
    "time": "20:00:00"
  }
  ```

#### **GET /reservations**
- **Descrição**: Lista todas as reservas.
- **Resposta**:
  ```json
  [
    {
      "id": 1,
      "restaurant_id": 1,
      "date": "2025-04-10",
      "time": "20:00:00"
    }
    {
      "id": 1,
      "restaurant_id": 1,
      "user_id": 1,
      "date": "2025-04-10",
      "time": "20:00:00"
    }
  ]

---

### **2. Restaurantes**
#### **GET /restaurants/:id**
- **Descrição**: Retorna os detalhes de um restaurante.
- **Resposta**:
  ```json
  {
      "id": 1,
      "name": "Restaurante A",
      "address": "Avenida Paulista 123",
      "date": "2025-04-10",
      "time": "20:00:00"
    }
  ```

#### **GET /restaurants**
- **Descrição**: Lista todos os restaurantes disponíveis.
- **Resposta**:
  ```json
  [
    {
      "id": 1,
      "name": "Restaurante A",
      "address": "Avenida Paulista 123",
      "date": "2025-04-10",
      "time": "20:00:00"
    },
    {
      "id": 2,
      "name": "Restaurante B",
      "address": "Avenida Paulista 234",
      "date": "2025-04-10",
      "time": "20:00:00"
    }
  ]
  ```

#### **POST /restaurants**
- **Descrição**: Adiciona um novo restaurante.
- **Body**:
  ```json
  {
    "name": "Restaurante C",
    "address": "Avenida Paulista 345"
  }
  ```

### **3. Users**
#### **GET /users/:id**
- **Descrição**: Retorna os detalhes de um usuário.
- **Resposta**:
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@mail.com"
  }
  ```

#### **GET /users**
- **Descrição**: Lista todos os usuários cadastrados.
- **Resposta**:
  ```json
  [
    {
      "id": 1,
      "name": "John Doe",
      "email": "john.doe@mail.com"
    },
    {
      "id": 2,
      "name": "Janne Doe",
      "email": "janne.doe@mail.com"
    },
  ]
  ```

#### **POST /users**
- **Descrição**: Adiciona um novo usuário.
- **Body**:
  ```json
  {
    "name": "Joe Doe",
    "email": "joe.doe@mail.com"
  }
  ```

### **4. Notifications**
#### **GET /notifications/:id**
- **Descrição**: Retorna os detalhes de uma notificação.
- **Resposta**:
  ```json
  {
		"id": 1,
		"user_id": 5,
		"message": "Welcome, John Doe!"
	},
  ```

#### **GET /notifications**
- **Descrição**: Lista todos as notificações cadastradas.
- **Resposta**:
  ```json
  [
    {
      "id": 1,
      "user_id": 1,
      "message": "Welcome, John Doe!"
    },
    {
      "id": 2,
      "user_id": 2,
      "message": "Welcome, Janne Doe!"
    },
  ]
  ```

#### **POST /notifications**
- **Descrição**: Adiciona uma nova notificação.
- **Body**:
  ```json
  {
    "user_id": 1,
    "message": "New Notification"
  }
  ```

---

## 📦 Comunicação entre Microsserviços

A comunicação entre os microsserviços é feita utilizando **RabbitMQ**. Os eventos principais incluem:

- **Reserva Criada**: Enviado quando uma nova reserva é criada.
- **Usuário criado**: Enviado para dar boas-vindas ao novo usuário criado.

---

## 🛠️ Como Executar o Projeto

1. Clone o repositório:
   ```bash
   git clone git@github.com:jonathanqueiroz/restaurants_reservations.git
   ```

2. Certifique-se de que o Docker e o Docker Compose estão instalados no seu sistema.

3. Suba os serviços utilizando o Docker Compose:

4. Inicie os microsserviços:
   ```bash
   docker-compose up
   ```

5. Certifique-se de que o RabbitMQ e os bancos de dados estão funcionando corretamente.

---
