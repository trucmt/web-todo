import TodoList from '../components/TodoList.js';
import TodoItem from '../components/TodoItem.js';

const todoList = new TodoList(document.getElementById('todo-list'));

document.getElementById('todo-form').addEventListener('submit', (e) => {
    e.preventDefault();
    const input = document.getElementById('todo-input');
    if (input.value.trim()) {
        todoList.addTodo(new TodoItem(input.value.trim()));
        input.value = '';
    }
});

// Tải danh sách công việc khi trang được tải
todoList.loadTodos();
