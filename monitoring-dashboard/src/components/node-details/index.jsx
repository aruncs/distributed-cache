import './style.scss'

export function NodeDetails(props) {
    const {id, name, data} = props
    return (
        <div className="node-details">
            <Tag tagKey={"Name"} value={name}/>
            <Tag tagKey="ID" value={id}/>
            <Tag tagKey="Data" value={<pre>{JSON.stringify(data, null, 2)}</pre>}/>
        </div>
    )
}


function Tag(props) {
    const {tagKey, value} = props
    return (
        <div className='tag'>
            <div className='tag-key'>{tagKey}</div>
            <div className='tag-value'>{value}</div>
        </div>
    )
}