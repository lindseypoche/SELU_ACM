// FUNCTIONAL COMPONENT
// import React from 'react'


// DESTUCTURING PROPS AND STATE 2
const Greet = (props) => {
    const {name, heroName, children} = props
    return(
        <div>
        <h1>Hello {name} a.k.a {heroName}</h1>
        {props.children}
        </div>
    )
    }
    
    export default Greet


// // DESTUCTURING PROPS AND STATE 1
// const Greet = ({name, heroName}) => {
//     return(
//         <div>
//         <h1>Hello {name} a.k.a {heroName}</h1>
//         </div>
//     )
//     }
    
//     export default Greet

// NEW SYNTAX
// const Greet = (props) => {
// console.log(props)
// return(
//     <div>
//     <h1>Hello {props.name} a.k.a {props.heroName}</h1>
//     {props.children}
//     </div>
// )
// }

// export default Greet


// CLASSIC FUNCTION
// import React from 'react'

// function Greet(props) {
//     return(
//         <div>
//         <h1> Hello {props.name} a.k.a {props.heroName}</h1>
//         {props.children}
//         </div>
//     )
// }

// export default Greet