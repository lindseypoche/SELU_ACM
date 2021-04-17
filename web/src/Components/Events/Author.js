import Avatar from './Avatar.js'

const Author = ({author}) => {

    return (
        <>
        {
            (author == null || author.avatar == null) ? (
                ''
            ) : (
                <Avatar avatar={author.avatar}/>
            )
        }
        </>
    )
}

export default Author