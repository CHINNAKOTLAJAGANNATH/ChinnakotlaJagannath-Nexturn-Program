document.addEventListener("DOMContentLoaded", () => {
    const taskInput = document.getElementById("task-input");
    const addTaskBtn = document.getElementById("add-task-btn");
    const taskList = document.getElementById("task-list");
    const pendingCount = document.getElementById("pending-count");
  
    const loadTasks = () => {
      const tasks = JSON.parse(localStorage.getItem("tasks")) || [];
      tasks.forEach(renderTask);
      updatePendingCount();
    };
  
    const saveTasks = () => {
      const tasks = Array.from(taskList.children).map(task => ({
        text: task.querySelector(".task-text").textContent,
        completed: task.classList.contains("completed"),
      }));
      localStorage.setItem("tasks", JSON.stringify(tasks));
    };
  
    const renderTask = ({ text, completed }) => {
      const task = document.createElement("li");
      task.className = `task ${completed ? "completed" : ""}`;
  
      const taskText = document.createElement("span");
      taskText.className = "task-text";
      taskText.textContent = text;
  
      const editBtn = document.createElement("button");
      editBtn.textContent = "Edit";
      editBtn.className = "edit-btn";
      editBtn.addEventListener("click", () => {
        const newText = prompt("Edit Task:", taskText.textContent);
        if (newText) {
          taskText.textContent = newText;
          saveTasks();
        }
      });
  
      const deleteBtn = document.createElement("button");
      deleteBtn.textContent = "Delete";
      deleteBtn.className = "delete-btn";
      deleteBtn.addEventListener("click", () => {
        task.remove();
        saveTasks();
        updatePendingCount();
      });
  
      task.appendChild(taskText);
      task.appendChild(editBtn);
      task.appendChild(deleteBtn);
  
      task.addEventListener("click", () => {
        task.classList.toggle("completed");
        saveTasks();
        updatePendingCount();
      });
  
      taskList.appendChild(task);
    };
  
    const updatePendingCount = () => {
      const pendingTasks = Array.from(taskList.children).filter(
        task => !task.classList.contains("completed")
      ).length;
      pendingCount.textContent = `Pending tasks: ${pendingTasks}`;
    };
  
    addTaskBtn.addEventListener("click", () => {
      const text = taskInput.value.trim();
      if (text) {
        renderTask({ text, completed: false });
        saveTasks();
        updatePendingCount();
        taskInput.value = "";
      }
    });
  
    loadTasks();
  });
  