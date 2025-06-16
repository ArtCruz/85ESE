import React from 'react';
import Table from 'react-bootstrap/Table';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';
import axios from 'axios';

class OrderList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            orders: [],
            product_id: '',
            quantity: '',
            products: [],
            showImageModal: false,
            modalImageUrl: '',
        };
    }

    componentDidMount() {
        this.fetchOrders();
        this.fetchProducts();
    }

    fetchOrders = () => {
        axios.get(window.global.gateway_location + '/orders')
            .then(res => this.setState({ orders: res.data }))
            .catch(err => console.error(err));
    }

    fetchProducts = () => {
        axios.get(window.global.gateway_location + '/products')
            .then(res => this.setState({ products: res.data }))
            .catch(err => console.error(err));
    }

    handleChange = (e) => {
        this.setState({ [e.target.name]: e.target.value });
    }

    handleSubmit = (e) => {
        e.preventDefault();
        const { product_id, quantity } = this.state;
        axios.post(window.global.gateway_location + '/orders', {
            product_id: parseInt(product_id),
            quantity: parseInt(quantity)
        })
        .then(() => {
            this.setState({ product_id: '', quantity: '' });
            this.fetchOrders();
        })
        .catch(err => alert("Erro ao criar ordem: " + err));
    }

    openImageModal = (imageUrl) => {
        this.setState({ showImageModal: true, modalImageUrl: imageUrl });
    }

    closeImageModal = () => {
        this.setState({ showImageModal: false, modalImageUrl: '' });
    }

    render() {
        return (
            <div style={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                minHeight: '100vh',
                background: 'linear-gradient(135deg, #a9744f 0%, #6f4e37 100%)',
                padding: '3rem'
            }}>
                <h1 style={{
                    marginBottom: '40px',
                    fontFamily: 'Pacifico, cursive',
                    color: '#fffbe7',
                    textShadow: '2px 2px 4px #6f4e37'
                }}>
                    Ordem de Compra
                </h1>
                <Form onSubmit={this.handleSubmit} style={{
                    background: '#fffbe7',
                    borderRadius: '16px',
                    boxShadow: '0 4px 24px rgba(111, 78, 55, 0.2)',
                    padding: '2rem',
                    marginBottom: '2rem',
                    minWidth: '320px',
                    maxWidth: '400px'
                }}>
                    <Form.Group>
                        <Form.Label>Produto</Form.Label>
                        <Form.Control as="select" name="product_id" value={this.state.product_id} onChange={this.handleChange} required>
                            <option value="">Selecione um produto</option>
                            {this.state.products.map(prod => (
                                <option key={prod.id} value={prod.id}>{prod.name}</option>
                            ))}
                        </Form.Control>
                    </Form.Group>
                    <Form.Group>
                        <Form.Label>Quantidade</Form.Label>
                        <Form.Control type="number" name="quantity" value={this.state.quantity} onChange={this.handleChange} min="1" required />
                    </Form.Group>
                    <Button type="submit" style={{ marginTop: '1rem', background: '#a9744f', border: 'none' }}>Criar Ordem</Button>
                </Form>
                <Table
                    striped
                    bordered
                    hover
                    responsive
                    style={{
                        backgroundColor: '#fffbe7',
                        borderRadius: '16px',
                        boxShadow: '0 4px 24px rgba(111, 78, 55, 0.2)',
                        overflow: 'hidden',
                        maxWidth: '800px'
                    }}
                >
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Produto</th>
                            <th>Imagem</th>
                            <th>Quantidade</th>
                        </tr>
                    </thead>
                    <tbody>
                        {this.state.orders.map(order => {
                            const prod = this.state.products.find(p => p.id === order.product_id);
                            const imageUrl = prod ? `${window.global.gateway_location}/images/${prod.id}` : '';
                            return (
                                <tr key={order.id}>
                                    <td>{order.id}</td>
                                    <td>{prod ? prod.name : order.product_id}</td>
                                    <td>
                                        {prod &&
                                            <img
                                                src={imageUrl}
                                                alt={prod.name}
                                                style={{
                                                    width: '50px',
                                                    height: '50px',
                                                    objectFit: 'cover',
                                                    cursor: 'pointer',
                                                    borderRadius: '50%',
                                                    border: '2px solid #a9744f',
                                                    background: '#fffbe7'
                                                }}
                                                onClick={() => this.openImageModal(imageUrl)}
                                                onError={e => { e.target.onerror = null; e.target.src = "https://via.placeholder.com/50"; }}
                                            />
                                        }
                                    </td>
                                    <td>{order.quantity}</td>
                                </tr>
                            );
                        })}
                    </tbody>
                </Table>

                <Modal show={this.state.showImageModal} onHide={this.closeImageModal} centered>
                    <Modal.Header closeButton style={{ background: '#3e2723', color: '#fffbe7' }}>
                        <Modal.Title>Imagem do Produto</Modal.Title>
                    </Modal.Header>
                    <Modal.Body style={{ background: '#3e2723', textAlign: 'center' }}>
                        <img
                            src={this.state.modalImageUrl}
                            alt="Produto"
                            style={{ maxWidth: '100%', maxHeight: '70vh', borderRadius: '16px', border: '4px solid #fffbe7' }}
                        />
                    </Modal.Body>
                </Modal>
            </div>
        );
    }
}

export default OrderList;