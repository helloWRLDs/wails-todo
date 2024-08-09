import './mainPage.css'
import { useEffect, useState } from 'react'
import { DeleteTodo, GetTodos, InsertTodo, ListTodos, UpdateTodo } from '../../../wailsjs/go/main/App'
import { formatDate, wrapPrior } from '../../format/formatters';
import { AddButton, DeleteButton, DoneButton, UndoneButton } from '../buttons/buttons';

const MainPage = () => {
    const [inputBody, setInputBody] = useState('')
    const [inputPrior, setInputPrior] = useState(4)
    const [error, setError] = useState('')


    const [todos, setTodos] = useState([])
    
    const loadTodos = () => {
        ListTodos()
            .then(elems => {
                if (elems && Array.isArray(elems)) {
                    setTodos(elems)
                } else {
                    setTodos([])
                }
                console.log(elems)
            }) 
            .catch(e => console.error(e))
    }

    const deleteTodo = (id) => {
        DeleteTodo(id)
            .then(resp => loadTodos())
            .catch(e => console.error(e))
    }

    const updateTodo = (todo) => {
        UpdateTodo(todo)
            .then(resp => loadTodos())
            .catch(e => console.error(e))
    }

    const addTodo = (e) => {
        e.preventDefault()
        if (inputBody.trim().length === 0) {
            setError("Todo body cannot be empty")
            return
        } else {
            setError('')
        }
        InsertTodo({Body: inputBody, Priority: inputPrior ?? 4})
            .then(resp => {loadTodos()})
            .catch(e => console.error(e)) 
    }

    useEffect(() => {
        loadTodos()
    }, [])

    return (
        <main>
            <form className='add-todo-form'>
                <h2 className='add-todo-title self-center'>Add new todo:</h2>
                <input type="text" name="Todo" className='self-center' id="todo" placeholder='Buy groceries..' onChange={(e) => setInputBody(e.target.value)}/>
                <select name="Priority" className='self-center' value={parseInt(inputPrior)} id="prior" onChange={(e) => setInputPrior(parseInt(e.target.value))}>
                    <option value={1}>Highest Priority</option>
                    <option value={2}>High Priority</option>
                    <option value={3}>Medium Priority</option>
                    <option value={4}>Low Priority</option>
                </select>
                <AddButton className="self-center" onClick={addTodo} />
            </form>
            <span className='err-span' style={{display: 'block'}}>{error}</span>

           {todos.length > 0 ? (
                <>
                <table className='todo-table'>
                    <thead>
                        <tr>
                            <th></th>
                            <th>Todo</th>
                            <th>Created At</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        {todos.map(elem => (
                            <tr key={elem.ID}>
                                <td>{wrapPrior(elem.Priority)}</td>
                                <td>
                                    {elem.IsDone ? (<span style={{textDecoration: 'line-through'}}>{elem.Body}</span>) : (<>{elem.Body}</>)}
                                </td>
                                <td>{formatDate(elem.CreatedAt)}</td>
                                <td style={{display: 'flex', margin: '0 auto'}}>
                                    < DeleteButton style={{marginRight: "15px"}} onClick={() => deleteTodo(elem.ID)} />
                                    {elem.IsDone ? 
                                    (
                                        < UndoneButton onClick={() => {
                                            const todo = elem
                                            elem.IsDone = false
                                            updateTodo(todo)
                                        }}/>
                                    ) 
                                    : 
                                    (
                                        < DoneButton onClick={() => {
                                            const todo = elem
                                            elem.IsDone = true
                                            updateTodo(todo)
                                        }}/>
                                    )}
                                </td>
                            </tr>
                        ))}
                    </tbody>
                </table>
                </>
           ) : (
                <>
                    <p>No Todos Yet</p>
                </>
           )}
        </main>
    )
}

export default MainPage;