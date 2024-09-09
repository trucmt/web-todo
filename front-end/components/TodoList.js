export default class TodoList {
    constructor(container) {
        this.container = container;
        this.todos = [];
    }

    addTodo(todoItem) {
        this.todos.push(todoItem);
        this.render();
        this.saveTodos();
    }

    removeTodo(id) {
        this.todos = this.todos.filter(todo => todo.id !== id);
        this.render();
        this.saveTodos();
    }

    render() {
        this.container.innerHTML = '';
        this.todos.forEach(todo => {
            const todoElement = todo.render();
            todoElement.querySelector('button').addEventListener('click', () => this.removeTodo(todo.id));
            this.container.appendChild(todoElement);
        });
    }

    saveTodos() {
        localStorage.setItem('todos', JSON.stringify(this.todos));
    }

    loadTodos() {
        const storedTodos = localStorage.getItem('todos');
        if (storedTodos) {
            this.todos = JSON.parse(storedTodos).map(todo => new TodoItem(todo.text, todo.id));
            this.render();
        }
    }
}
