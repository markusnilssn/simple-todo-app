const API = "http://backend:8080";

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
                <button onclick='openEditModal(${JSON.stringify(todo)})'>Edit</button>
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

function openEditModal(todo) {
    document.getElementById("editId").value = todo.id;
    document.getElementById("editTitle").value = todo.title;
    document.getElementById("editDescription").value = todo.description;
    document.getElementById("editPriority").value = todo.priority;

    document.getElementById("editModal").style.display = "flex";
}

function closeModal() {
    document.getElementById("editModal").style.display = "none";
}

async function handleUpdateTodo() {
    try {
        const id = parseInt(document.getElementById("editId").value);
        const title = document.getElementById("editTitle").value;
        const description = document.getElementById("editDescription").value;
        const priority = parseInt(document.getElementById("editPriority").value);

        await updateTodo(id, { title, description, priority });
        closeModal();
        await loadAllTodos();
    } catch (err) {
        closeModal();
        handleError(err, "Failed to update todo");
    }
}

/* Priority Helpers */
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

async function loadTodos() {
    if (verbose)
        console.log("load Todos")

    const res = await fetch(`${API}/todos`);
    const todos = await res.json();

    const list = document.getElementById("todoList");
    list.innerHTML = "";

    todos.forEach(todo => {
        const li = document.createElement("li");
        li.innerText = todo.title;
        list.appendChild(li);
    });
}

function handleError(error, fallbackMessage = "Something went wrong") {
    if (verbose)
        console.warn(error);

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

async function addTodo() {
    if (verbose)
        console.log("add Todo");

    const input = document.getElementById("todoInput");

    await fetch(`${API}/todos`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title: input.value })
    });

    input.value = "";
    loadTodos();
}

async function removeTodo() {
    if (verbose)
        console.log("remove Todo")
}
async function updateTodo() {
    if (verbose)
        console.log("update Todo");
}

