// API Configuration
const API_BASE_URL = 'http://localhost:8080/api';

// Load todos when page loads
document.addEventListener('DOMContentLoaded', () => {
    fetchTodos();
});

// Fetch all todos
async function fetchTodos() {
    try {
        const response = await fetch(`${API_BASE_URL}/todos`);
        const todos = await response.json();
        displayTodos(todos);
    } catch (error) {
        console.error('Error fetching todos:', error);
    }
}

// Display todos in the UI
function displayTodos(todos) {
    const todoList = document.getElementById('todoList');
    todoList.innerHTML = '';
    
    todos.forEach(todo => {
        const todoElement = createTodoElement(todo);
        todoList.appendChild(todoElement);
    });
}

// Create a todo element
function createTodoElement(todo) {
    const div = document.createElement('div');
    div.className = `todo-item ${todo.completed ? 'completed' : ''}`;
    div.dataset.id = todo.id;
    
    div.innerHTML = `
        <input type="checkbox" class="todo-checkbox" ${todo.completed ? 'checked' : ''} onchange="toggleTodo('${todo.id}')">
        <span class="todo-title">${todo.title}</span>
        <div class="todo-actions">
            <button class="edit-btn" onclick="editTodo('${todo.id}')">Edit</button>
            <button class="delete-btn" onclick="deleteTodo('${todo.id}')">Delete</button>
        </div>
    `;
    
    return div;
}

// Add new todo
async function addTodo() {
    const input = document.getElementById('todoInput');
    const title = input.value.trim();
    
    if (!title) return;
    
    try {
        const response = await fetch(`${API_BASE_URL}/todos`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                title: title,
                completed: false
            })
        });
        
        if (response.ok) {
            input.value = '';
            fetchTodos(); // Refresh the list
        }
    } catch (error) {
        console.error('Error adding todo:', error);
    }
}

// Toggle todo completion status
async function toggleTodo(id) {
    try {
        // First get the current todo
        const getResponse = await fetch(`${API_BASE_URL}/todos/${id}`);
        const todo = await getResponse.json();
        
        // Then update it
        const updateResponse = await fetch(`${API_BASE_URL}/todos/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                ...todo,
                completed: !todo.completed
            })
        });
        
        if (updateResponse.ok) {
            fetchTodos(); // Refresh the list
        }
    } catch (error) {
        console.error('Error toggling todo:', error);
    }
}

// Edit todo
async function editTodo(id) {
    const newTitle = prompt('Enter new title:');
    if (!newTitle) return;
    
    try {
        // First get the current todo
        const getResponse = await fetch(`${API_BASE_URL}/todos/${id}`);
        const todo = await getResponse.json();
        
        // Then update it
        const updateResponse = await fetch(`${API_BASE_URL}/todos/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                ...todo,
                title: newTitle
            })
        });
        
        if (updateResponse.ok) {
            fetchTodos(); // Refresh the list
        }
    } catch (error) {
        console.error('Error editing todo:', error);
    }
}

// Delete todo
async function deleteTodo(id) {
    if (!confirm('Are you sure you want to delete this todo?')) return;
    
    try {
        const response = await fetch(`${API_BASE_URL}/todos/${id}`, {
            method: 'DELETE'
        });
        
        if (response.ok) {
            fetchTodos(); // Refresh the list
        }
    } catch (error) {
        console.error('Error deleting todo:', error);
    }
}