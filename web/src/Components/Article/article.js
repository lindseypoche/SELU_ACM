import './Article.css'

const Article = ({ event }) => {

      return (
        <article className="article">
          {/* <h3 className="article__category" style={styles}>{event.category}</h3> */}
          <h2 className="article__title">{event.content}</h2>
          {/* <p className="article__excerpt">{details.excerpt}</p> */}
        </article>
      )
}

export default Article 