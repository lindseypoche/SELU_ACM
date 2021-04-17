
const Avatar = ({avatar}) => {

    console.log("AVATAR:>>>", avatar)

    return (
        <img className="avatar__image" src={avatar.image_url} />
    )
}

export default Avatar;