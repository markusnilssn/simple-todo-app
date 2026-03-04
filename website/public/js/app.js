const API = "http://localhost:8080";

var verbose = true;

document.addEventListener("DOMContentLoaded", () => {
    loadAllTodos();
});

async function loadAllTodos() {
    try {
        const todos = await loadTodos();
        renderTodos(todos);
    } catch (err) {
        handleError(err, "Failed to load todos");
    }
}

function renderTodos(todos) {
    const list = document.getElementById("todoList");
    list.innerHTML = "";

    if (!todos || todos.length === 0) {
        showEmptyState();
        return;
    }

    hideEmptyState();

    todos.forEach(todo => {
        const div = document.createElement("div");
        div.className = "todo-item";

        div.innerHTML = `
            <div class="todo-info">
                <strong>${todo.title}</strong><br>
                <small>${todo.description}</small><br>
                <span class="priority ${getPriorityClass(todo.priority)}">
                    ${getPriorityText(todo.priority)}
                </span>
            </div>
            <div class="todo-actions">
                <button onclick="handleDelete(${todo.id})">Delete</button>
            </div>
        `;

        list.appendChild(div);
    });
}

async function handleAddTodo() {
    try {
        const title = document.getElementById("titleInput").value;
        const description = document.getElementById("descriptionInput").value;
        const priority = parseInt(document.getElementById("priorityInput").value);

        await addTodo({ title, description, priority });
        await loadAllTodos();
    } catch (err) {
        handleError(err, "Failed to add todo");
    }
}

async function handleDelete(id) {
    try {
        await removeTodo(id);
        await loadAllTodos();
    } catch (err) {
        handleError(err, "Failed to delete todo");
    }
}

async function loadTodos() {
    if (verbose) console.log("Fetching todos from API...");

    const res = await fetch(`${API}/todos`);

    if (!res.ok) throw new Error("Could not fetch todos");

    const data = await res.json();

    console.log(data);
    return data;
}

function handleError(error, fallbackMessage = "Something went wrong") {
    if (verbose) console.warn(error);

    clearTodos();
    showEmptyState();
    showToast(error?.message || fallbackMessage);
}

function clearTodos() {
    const list = document.getElementById("todoList");
    list.innerHTML = "";
}

function showEmptyState() {
    document.getElementById("emptyState").style.display = "block";
}

function hideEmptyState() {
    document.getElementById("emptyState").style.display = "none";
}

async function addTodo(todoData) {
    if (verbose) console.log("add Todo:", todoData);

    const res = await fetch(`${API}/todos`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(todoData)
    });

    if (!res.ok) {
        throw new Error(`Server error: ${res.status}`);
    }

    document.getElementById("titleInput").value = "";
    document.getElementById("descriptionInput").value = "";
}

async function removeTodo(id) {
    if (verbose) console.log("remove Todo ID:", id);

    const res = await fetch(`${API}/todos/${id}`, {
        method: "DELETE"
    });

    if (!res.ok) {
        throw new Error(`Failed to delete task ${id}`);
    }
}

async function updateTodo(id, updatedData) {
    if (verbose) console.log("update Todo ID:", id, updatedData);

    const res = await fetch(`${API}/todos/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(updatedData)
    });

    if (!res.ok) {
        throw new Error(`Failed to update task ${id}`);
    }
}

function getPriorityText(p) {
    return ["Low", "Medium", "High", "Critical"][p];
}

function getPriorityClass(p) {
    return ["low", "medium", "high", "critical"][p];
}

/* Toast */
function showToast(message) {
    const container = document.getElementById("toastContainer");

    const toast = document.createElement("div");
    toast.className = "toast";
    toast.innerText = message;

    container.appendChild(toast);

    setTimeout(() => {
        toast.remove();
    }, 3300);
}