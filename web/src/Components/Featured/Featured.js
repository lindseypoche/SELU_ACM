import './Featured.css'

const Featured = ({ featured }) => {
 
    return (
        <>
        <div className='featured'>
              { featured.message.attachments != null ? 
                    (
                        <>
                        <p>{featured.channel_id}</p>
                        <img src={featured.message.attachments.url} />
                        <div className='content'>
                            <p>{featured.message.content}</p>
                        </div>
                        </>
                    ) : (
                        <p>
                            No featured event found
                        </p>
                    )
              }
        </div>
        < br/><br/><br/>
        </>
    )
}

export default Featured
