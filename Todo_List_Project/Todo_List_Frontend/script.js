document.addEventListener("DOMContentLoaded", function () {
  const todoListContainer = document.getElementById("todo-list");
  const todoForm = document.getElementById("todo-form");
  const todoInput = document.getElementById("todo-input");
  const isCompletedInput = document.getElementById("is-completed");

  const updateForm = document.getElementById("update-form");
  const updateTodoInput = document.getElementById("update-todo-input");
  const updateIsCompletedInput = document.getElementById("update-is-completed");
  const updateTodoId = document.getElementById("update-todo-id");

  // Fetch todo lists from API
  function fetchTodos() {
    fetch("http://localhost:8900/api/getAllTodoLists")
      .then((response) => response.json())
      .then((data) => {
        if (data.code === 200) {
          const todoLists = data.data;

          // Clear the container before rendering
          todoListContainer.innerHTML = "";

          // Render each todo item
          todoLists.forEach((todo) => {
            const todoItem = document.createElement("div");
            todoItem.classList.add("todo-item");

            if (todo.is_completed) {
              todoItem.classList.add("completed");
            }

            todoItem.innerHTML = `
              <span class="todo-text">${todo.todo}</span>
              <span>${new Date(todo.created_at).toLocaleString()}</span>
              <button class="edit-button" data-id="${todo.id}">Edit</button>
              <button class="delete-button" data-id="${todo.id}">Delete</button>
            `;

            todoListContainer.appendChild(todoItem);
          });

          // Add event listeners to all edit buttons
          document.querySelectorAll(".edit-button").forEach((button) => {
            button.addEventListener("click", function () {
              const todoId = this.getAttribute("data-id");
              const todoText =
                this.parentElement.querySelector(".todo-text").textContent; // Get the current todo text
              const isCompleted =
                this.parentElement.classList.contains("completed"); // Get the completion status
              showUpdateForm(todoId, todoText, isCompleted); // Pass data to the form
            });
          });
          // Add event listeners to all delete buttons
          document.querySelectorAll(".delete-button").forEach((button) => {
            button.addEventListener("click", function () {
              const todoId = this.getAttribute("data-id");
              deleteTodo(todoId);
            });
          });
        } else {
          todoListContainer.innerHTML = `<p>Failed to load todo lists. Please try again later.</p>`;
        }
      })
      .catch((error) => {
        console.error("Error fetching data:", error);
        todoListContainer.innerHTML = `<p>Error fetching data. Please try again later.</p>`;
      });
  }

  // Function to delete a todo
  function deleteTodo(id) {
    fetch(`http://localhost:8900/api/deleteTodoList/${id}`, {
      method: "DELETE",
      headers: {
        "Content-Type": "application/json",
      },
    })
      .then((response) => response.json())
      .then((result) => {
        if (result.code === 200) {
          alert("Todo deleted successfully!");
          fetchTodos(); // Refresh the todo list after deletion
        } else {
          alert("Failed to delete todo. Please try again.");
        }
      })
      .catch((error) => {
        console.error("Error deleting todo:", error);
        alert("Error deleting todo. Please try again later.");
      });
  }

  // Function to show the update form
  function showUpdateForm(todoId, todoText, isCompleted) {
    updateTodoId.value = todoId; // Set the ID of the todo to be updated
    updateTodoInput.value = todoText; // Set the current text of the todo
    updateIsCompletedInput.checked = isCompleted; // Set the current completed status
    updateForm.style.display = "block"; // Show the update form
  }

  // Function to update a todo
  function updateTodo(id, todo, isCompleted) {
    const data = {
      todo: todo,
      is_completed: isCompleted,
    };

    fetch(`http://localhost:8900/api/updateTodoList/${id}`, {
      method: "PUT",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((result) => {
        if (result.code === 200) {
          alert("Todo updated successfully!");
          fetchTodos(); // Refresh the todo list after updating
          updateForm.style.display = "none"; // Hide the update form
          updateTodoInput.value = ""; // Clear the input field
          updateIsCompletedInput.checked = false; // Reset the checkbox
        } else {
          alert("Failed to update todo. Please try again.");
        }
      })
      .catch((error) => {
        console.error("Error updating todo:", error);
        alert("Error updating todo. Please try again later.");
      });
  }

  // Function to create a new todo
  function createTodo(todo, isCompleted) {
    const data = {
      todo: todo,
      is_completed: isCompleted,
    };

    fetch("http://localhost:8900/api/createTodoList", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    })
      .then((response) => response.json())
      .then((result) => {
        if (result.code === 200) {
          alert("Todo created successfully!");
          fetchTodos(); // Refresh the todo list after creating
          todoInput.value = ""; // Clear the input field
          isCompletedInput.checked = false; // Reset the checkbox
        } else {
          alert("Failed to create todo. Please try again.");
        }
      })
      .catch((error) => {
        console.error("Error creating todo:", error);
        alert("Error creating todo. Please try again later.");
      });
  }

  // Handle form submission for updating todos
  updateForm.addEventListener("submit", function (event) {
    event.preventDefault();

    const id = updateTodoId.value.trim();
    const updatedTodo = updateTodoInput.value.trim();
    const isCompleted = updateIsCompletedInput.checked;

    if (updatedTodo && id) {
      updateTodo(id, updatedTodo, isCompleted);
    } else {
      alert("Please enter a valid todo and ensure the ID is present.");
    }
  });

  // Handle form submission for creating todos
  todoForm.addEventListener("submit", function (event) {
    event.preventDefault();

    const newTodo = todoInput.value.trim();
    const isCompleted = isCompletedInput.checked;

    if (newTodo) {
      createTodo(newTodo, isCompleted);
    } else {
      alert("Please enter a valid todo.");
    }
  });

  // Initial fetch of todos on page load
  fetchTodos();
});
