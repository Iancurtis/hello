import React from 'react';
import ReactDom from 'react-dom';

class CommentBox extends React.Component {
    render() {
        return (
            <div>
                <h2>Hello, I'm a comment box.</h2>
            </div>  
        );
    }
}

ReactDom.render(<CommentBox />, document.getElementById('content'))