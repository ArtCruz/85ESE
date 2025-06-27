import React from 'react';
import Table from 'react-bootstrap/Table';
import Modal from 'react-bootstrap/Modal';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

function ImagemButton({ productId }) {
    const navigate = useNavigate();
    return (
        <button
            className="btn btn-info"
            style={{ background: '#6f4e37', border: 'none', borderRadius: '50%', width: '36px', height: '36px', padding: 0, fontSize: '1.2rem' }}
            title="Upload de Imagem"
            onClick={() => navigate(`/upload-image/${productId}`)}
        >
            📷
        </button>
    );
}

class ProductList extends React.Component {
    constructor(props) {
        super(props);
        this.state = {
            products: [],
            showAddProductDialog: false,
            isEditing: false,
            newProduct: {
                id: '',
                name: '',
                price: '',
                sku: '',
            },
            showImageModal: false,
            modalImageUrl: '',
        };

        this.readData = this.readData.bind(this);
        this.showAddProductDialog = this.showAddProductDialog.bind(this);
        this.closeAddProductDialog = this.closeAddProductDialog.bind(this);
        this.handleInputChange = this.handleInputChange.bind(this);
        this.handleSaveProduct = this.handleSaveProduct.bind(this);
    }

    componentDidMount() {
        this.readData();
    }
// testes
    readData() {
        axios.get(window.global.gateway_location + '/products')
            .then((response) => {
                this.setState({ products: response.data });
            })
            .catch((error) => {
                console.log(error);
            });
    }

    getProducts() {
        let table = [];

        for (let i = 0; i < this.state.products.length; i++) {
            const product = this.state.products[i];
            table.push(
                <tr key={i}>
                    <td>{product.id}</td>
                    <td>{product.name}</td>
                    <td>{product.description}</td>
                    <td>{product.price}</td>
                    <td>{product.sku}</td>
                    <td>
                        <img
                            src={`${window.global.gateway_location}/images/${product.id}`}
                            alt="Produto"
                            style={{
                                width: '50px',
                                height: '50px',
                                objectFit: 'cover',
                                cursor: 'pointer',
                                borderRadius: '50%',
                                border: '2px solid #a9744f',
                                background: '#fffbe7'
                            }}
                            onClick={() => this.openImageModal(`${window.global.gateway_location}/images/${product.id}`)}
                            onError={e => { e.target.src = "https://via.placeholder.com/50"; }}
                        />
                    </td>
                    <td>
                        <button
                            className="btn btn-secondary"
                            style={{ marginRight: '0.5rem', background: '#a9744f', border: 'none' }}
                            onClick={() => this.editProduct(product)}
                        >
                            Edit
                        </button>
                        <button
                            className="btn btn-danger"
                            style={{ marginRight: '0.5rem', background: '#c97c5d', border: 'none' }}
                            onClick={() => this.deletarProduto(product.id)}
                        >
                            Delete
                        </button>
                        <ImagemButton productId={product.id} />
                    </td>
                </tr>
            );
        }

        return table;
    }

    deletarProduto(productId) {
        if (window.confirm("Tem certeza que deseja deletar este produto?")) {
            axios.delete(window.global.gateway_location + `/products/${productId}`)
                .then((response) => {
                    this.readData(); // Atualiza a lista de produtos
                })
                .catch((error) => {
                    console.error("Erro ao deletar produto:", error);
                });
        }
    }

    showAddProductDialog() {
        this.setState({ showAddProductDialog: true });
    }

    closeAddProductDialog() {
        this.setState({
            showAddProductDialog: false,
            isEditing: false,
            newProduct: {
                id: '',
                name: '',
                price: '',
                sku: '',
            },
        });
    }

    handleInputChange(event) {
        const { name, value } = event.target;
        this.setState((prevState) => ({
            newProduct: {
                ...prevState.newProduct,
                [name]: value,
            },
        }));
    }

    handleSaveProduct() {
        const { newProduct, isEditing } = this.state;

        if (isEditing) {
            axios.put(window.global.gateway_location + `/products/${newProduct.id}`, newProduct)
                .then((response) => {
                    const data = response.data;

                    if (data.messages && Array.isArray(data.messages)) {
                        const errorMessages = data.messages;

                        if (errorMessages.some(msg => msg.includes("Product.SKU") && msg.includes("validation for 'SKU' failed"))) {
                            alert("Formatação errada de SKU");
                        } else {
                            alert("Erro desconhecido:\n" + errorMessages.join('\n'));
                        }

                        console.error('Erro ao salvar produto:', errorMessages);
                        return; // Impede a continuação
                    }

                    this.closeAddProductDialog();
                    this.readData();
                })
        } else {
            const { id, price, ...rest } = this.state.newProduct;
            const productToSend = {
                ...rest,
                price: parseFloat(price)
            };

            axios.post(window.global.gateway_location + '/products', productToSend)
                .then((response) => {
                    this.closeAddProductDialog();
                    this.readData();
                })
                .catch((error) => {
                    if (error.response && error.response.data && error.response.data.Messages) {
                        alert(error.response.data.Messages.join('\n'));
                    } else if (error.response && error.response.data && error.response.data.message) {
                        alert(error.response.data.message);
                    } else {
                        alert('Formatação errada do SKU');
                    }
                    console.error('Erro ao adicionar produto:', error);
                });
        }
    }

    editProduct(product) {
        this.setState({
            newProduct: { ...product },
            showAddProductDialog: true,
            isEditing: true,
        });
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
                    Coffee Shop Menu
                </h1>
                <Table
                    striped
                    bordered
                    hover
                    responsive
                    style={{
                        backgroundColor: '#fffbe7',
                        borderRadius: '16px',
                        boxShadow: '0 4px 24px rgba(111, 78, 55, 0.2)',
                        overflow: 'hidden'
                    }}
                >
                    <thead>
                        <tr>
                            <th>ID</th>
                            <th>Name</th>
                            <th>Description</th>
                            <th>Price</th>
                            <th>SKU</th>
                            <th>Image</th>
                            <th>Options</th>
                        </tr>
                    </thead>
                    <tbody>
                        {this.getProducts()}
                    </tbody>
                    <tfoot>
                        <tr>
                            <td colSpan="7">
                                <button
                                    className="btn btn-primary"
                                    style={{ background: '#a9744f', border: 'none' }}
                                    onClick={this.showAddProductDialog}
                                >
                                    Add Product
                                </button>
                            </td>
                        </tr>
                    </tfoot>
                </Table>

                <Modal
                    show={this.state.showAddProductDialog}
                    onHide={this.closeAddProductDialog}
                >
                    <Modal.Header closeButton>
                        <Modal.Title>
                            {this.state.isEditing ? 'Edit Product' : 'Add Product'}
                        </Modal.Title>
                    </Modal.Header>
                    <Modal.Body>
                        <Form>
                            <Form.Group controlId="productName">
                                <Form.Label>Name</Form.Label>
                                <Form.Control
                                    type="text"
                                    name="name"
                                    placeholder="Enter product name"
                                    value={this.state.newProduct.name}
                                    onChange={this.handleInputChange}
                                />
                            </Form.Group>
                            <Form.Group controlId="productDescription">
                                <Form.Label>Description</Form.Label>
                                <Form.Control
                                    type="text"
                                    name="description"
                                    placeholder="Enter product description"
                                    value={this.state.newProduct.description}
                                    onChange={this.handleInputChange}
                                />
                            </Form.Group>
                            <Form.Group controlId="productPrice">
                                <Form.Label>Price</Form.Label>
                                <Form.Control
                                    type="number"
                                    name="price"
                                    placeholder="Enter product price"
                                    value={this.state.newProduct.price}
                                    onChange={this.handleInputChange}
                                />
                            </Form.Group>
                            <Form.Group controlId="productSKU">
                                <Form.Label>SKU</Form.Label>
                                <Form.Control
                                    type="text"
                                    name="sku"
                                    placeholder="Enter product SKU"
                                    value={this.state.newProduct.sku}
                                    onChange={this.handleInputChange}
                                />
                            </Form.Group>
                        </Form>
                    </Modal.Body>
                    <Modal.Footer>
                        <Button variant="secondary" onClick={this.closeAddProductDialog}>
                            Close
                        </Button>
                        <Button variant="primary" onClick={this.handleSaveProduct}>
                            Save Changes
                        </Button>
                    </Modal.Footer>
                </Modal>

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

export default ProductList;