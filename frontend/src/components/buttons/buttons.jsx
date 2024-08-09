import './buttons.css'
import { MdDeleteForever, MdOutlineDone } from "react-icons/md";
import { MdOutlineRemoveDone } from "react-icons/md";
import { IoMdDoneAll } from "react-icons/io";
import { FaPlus } from "react-icons/fa6";

export const DeleteButton = (props) => {
    return (
        <button {...props} className='delete-btn'>
            <MdDeleteForever />
        </button>
    )
}

export const DoneButton = (props) => {
    return (
        <button {...props} className='done-btn'>
            {/* <MdOutlineDone /> */}
            <IoMdDoneAll />
        </button>
    )
}

export const UndoneButton = (props) => {
    return (
        <button {...props} className='undone-btn'>
            < MdOutlineRemoveDone/>
        </button>
    )
}

export const AddButton = (props) => {
    return (
        <button {...props} className='add-btn'>
            <FaPlus />
        </button>
    )
}