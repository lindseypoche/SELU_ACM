import './Featured.css'

const Featured = ({ message, attachment }) => {
 
    return (
        <>
        <div className='featured'>
            <img src={attachment.url} />
            <div className='content'>
                <p>{message.content}</p>
            </div>
        </div>
        < br/><br/><br/>
        </>
    )
}

export default Featured
