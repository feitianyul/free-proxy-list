**中文** | [English (README_EN.md)](README_EN.md)

---

<p align="center">
  <img src="https://img.shields.io/badge/每1小时更新-通过-success">  
  <br>
  <img src="https://img.shields.io/website/https/getfreeproxy.com.svg">
  <img src="https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/total.svg">
  <img src="https://img.shields.io/github/last-commit/feitianyul/free-proxy-list.svg">
  <img src="https://img.shields.io/github/license/feitianyul/free-proxy-list.svg">
  
  <br>
  <br>
  <a href="https://getfreeproxy.com/lists/" title="可用代理列表">可用代理列表</a> | <a href="https://getfreeproxy.com/tools/proxy-checker" title="在线代理检测">免费代理检测</a> | <a href="https://getfreeproxy.com/tools/proxy-protocol-parser" title="代理协议解析">通用代理协议解析</a> | <a href="https://developer.getfreeproxy.com/" title="代理 API">免费代理 API</a>
  <br>
</p>

# 🌎 GetFreeProxy (GFP)：免费代理列表

**GetFreeProxy (GFP)** 是一个开源项目，自动从互联网聚合并校验免费代理，旨在为开发者、研究人员及需要代理服务的用户提供新鲜、可靠、可用的公共代理列表。

列表按小时更新，确保您始终能获取到最新的可用代理。

---

## 📖 项目说明

本项目为开源免费代理聚合与校验工具，从互联网公开源拉取代理并**仅保留 HTTP、HTTPS** 两种类型，经校验后生成列表，供开发者、研究人员等使用。

### 本仓库特点

- **仅保留两种代理**：HTTP、HTTPS，不收录 SOCKS、VMess、Trojan、VLESS、SS/SSR、Hysteria 等其它协议。
- **校验规则**：五域名中**任意 3 个**在 2 秒内成功（HTTP 200）即视为该协议通过。对每条代理访问以下五个地址验证（优先 HEAD，不支持则回退 GET）：
  - `https://www.eastmoney.com/`
  - `https://www.sse.com.cn/`
  - `https://finance.sina.com.cn/`（新浪财经）
  - `https://web.ifzq.gtimg.cn/`
  - `https://proxy.finance.qq.com/`
  每个代理分别以 **HTTP 代理** 和 **HTTPS 代理** 各测一次；**协议** 写入 meta：只通 HTTP→`http`，只通 HTTPS→`https`，两个都通→`http/s`。去重按 **协议+IP+端口**。校验时**多代理并发**、**单代理内五域名并行**。列表直显：表格列为「代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议」。
- **更新频率**：列表按小时更新，保证可用代理的时效性。
- **并发参数**：校验 worker 数可通过 `-check-workers`（如 `-check-workers=4000`）或环境变量 `GFP_CHECK_WORKERS` 设置，默认 4000，最大 4000。遇目标站限流可适当调低。

### 工作流程

1. **拉取**：从 `sources/` 目录下配置的源（仅处理 `http.txt`、`https.txt`）拉取原始代理数据，支持动态 URL 及 Base64 等格式。
2. **解析与规范化**：将原始数据解析为标准代理格式（协议、IP、端口、认证等）。
3. **校验**：对 HTTP/HTTPS 代理通过上述验证与 2 秒超时规则进行筛选。
4. **去重与存储**：通过校验的代理去重后写入内存。
5. **生成列表**：按协议生成 `list/` 目录下的 `http.txt`、`https.txt`，并更新统计与 README 中的下载表格。

自动化由 GitHub Actions 执行：**全量流程**（抓取→解析→验证→生成列表）**每 6 小时**运行一次；**轻量复测**（对已有列表做连通性复测、剔除失效代理）**每 1 小时**运行一次。全量任务最长运行 12 小时，超时才会取消。下表「最后更新」时间为 UTC 及 UTC+8。

### 支持的代理格式示例

| 类型 | 格式 | 示例 |
| :--- | :--- | :--- |
| **HTTP/HTTPS** | `http://ip:port` | `http://1.2.3.4:8080` |
| | `https://ip:port` | `https://1.2.3.4:8080` |
| | `http://user:pass@ip:port` | `http://user:pass@1.2.3.4:8080` |

---

## 🔗 直接下载链接

点击下方表格中您需要的协议类型即可获取最新列表，链接始终指向最近更新的代理文件。上述三个文件（http.txt、https.txt、passed.txt）同时发布至 [GitHub Releases (tag: lists)](https://github.com/feitianyul/free-proxy-list/releases/tag/lists)，每次运行覆盖同版本附件，可固定使用该 Release 的附件 URL。

<!-- BEGIN PROXY LIST -->

最后更新：2026-04-12 17:33:40 UTC（2026-04-13 01:33:40 UTC+8）

**代理总数：53**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 53 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 218.108.131.186:17890 | ✓ 1172ms | ✓ 1934ms | ✓ 1045ms | ✓ 1284ms | ✓ 1069ms | http |
| 1.231.81.166:3128 | ✓ 1664ms | ✓ 1265ms | 否 | ✓ 1391ms | ✓ 1021ms | http |
| 147.161.210.140:8800 | ✓ 1638ms | 否 | ✓ 911ms | ✓ 1912ms | 否 | http |
| 113.160.132.26:8080 | 否 | ✓ 1582ms | ✓ 1449ms | ✓ 1436ms | ✓ 1155ms | http |
| 45.167.124.52:8080 | ✓ 721ms | 否 | ✓ 1246ms | 否 | ✓ 1640ms | http |
| 35.225.22.61:80 | ✓ 1060ms | 否 | ✓ 498ms | ✓ 1140ms | 否 | http |
| 223.84.151.86:30005 | ✓ 1585ms | ✓ 1508ms | ✓ 1224ms | ✓ 1665ms | ✓ 1660ms | http |
| 167.103.115.102:8800 | 否 | 否 | ✓ 1361ms | ✓ 1321ms | ✓ 1223ms | http |
| 167.103.144.127:8800 | ✓ 1783ms | 否 | ✓ 1432ms | 否 | ✓ 1694ms | http |
| 217.76.245.80:999 | ✓ 710ms | 否 | ✓ 1168ms | ✓ 1544ms | ✓ 1087ms | http |
| 45.167.125.21:999 | ✓ 1030ms | ✓ 1919ms | ✓ 1081ms | ✓ 1947ms | ✓ 1453ms | http |
| 5.196.101.18:3128 | ✓ 872ms | 否 | ✓ 948ms | ✓ 1808ms | ✓ 1480ms | http |
| 120.92.108.86:7890 | ✓ 1392ms | 否 | 否 | ✓ 1950ms | ✓ 1629ms | http |
| 167.103.34.108:8800 | ✓ 1727ms | 否 | ✓ 1694ms | ✓ 1499ms | 否 | http |
| 147.161.239.240:8800 | ✓ 709ms | ✓ 1468ms | ✓ 1137ms | ✓ 1358ms | ✓ 1280ms | http |
| 8.219.195.129:1080 | ✓ 1021ms | ✓ 1955ms | ✓ 926ms | ✓ 1273ms | ✓ 1001ms | http |
| 8.219.97.248:80 | ✓ 1399ms | 否 | ✓ 1605ms | 否 | ✓ 1478ms | http |
| 5.104.87.17:8051 | ✓ 1929ms | 否 | ✓ 1697ms | 否 | ✓ 1648ms | http |
| 162.240.154.26:3128 | ✓ 797ms | ✓ 1920ms | ✓ 1728ms | 否 | 否 | http |
| 116.80.60.44:7777 | ✓ 1660ms | 否 | ✓ 1668ms | 否 | ✓ 1845ms | http |
| 177.234.217.88:999 | ✓ 1448ms | ✓ 1715ms | ✓ 1024ms | ✓ 1666ms | ✓ 1334ms | http |
| 45.140.147.82:1081 | ✓ 549ms | ✓ 1655ms | ✓ 668ms | ✓ 1332ms | ✓ 1152ms | http |
| 45.140.147.82:1082 | ✓ 1489ms | ✓ 1199ms | ✓ 390ms | ✓ 1870ms | ✓ 1293ms | http |
| 212.58.132.5:8888 | ✓ 1114ms | 否 | 否 | ✓ 1525ms | ✓ 1164ms | http |
| 103.157.200.126:3128 | ✓ 1449ms | 否 | 否 | ✓ 1540ms | ✓ 1212ms | http |
| 46.30.46.133:3128 | ✓ 1307ms | ✓ 1365ms | ✓ 1862ms | ✓ 1866ms | ✓ 1709ms | http |
| 34.101.184.164:3128 | ✓ 1844ms | 否 | ✓ 1465ms | ✓ 1766ms | ✓ 1252ms | http |
| 167.103.31.122:8800 | ✓ 1595ms | 否 | 否 | ✓ 1686ms | ✓ 1794ms | http |
| 101.32.244.83:8080 | ✓ 1351ms | ✓ 1736ms | ✓ 1136ms | ✓ 1732ms | ✓ 1510ms | http |
| 121.43.196.210:8222 | ✓ 1129ms | ✓ 1284ms | ✓ 989ms | ✓ 1348ms | ✓ 1127ms | http |
| 121.43.196.213:8222 | ✓ 1111ms | ✓ 1284ms | ✓ 1021ms | ✓ 1421ms | ✓ 1068ms | http |
| 114.55.226.123:10086 | ✓ 1990ms | ✓ 1729ms | ✓ 1231ms | ✓ 1504ms | ✓ 1235ms | http |
| 36.141.21.200:7890 | ✓ 1108ms | ✓ 1339ms | ✓ 1108ms | ✓ 1435ms | ✓ 1167ms | http |
| 5.255.123.43:1080 | ✓ 1020ms | ✓ 1437ms | ✓ 361ms | 否 | 否 | http |
| 207.254.71.62:8088 | ✓ 955ms | ✓ 1642ms | ✓ 1259ms | 否 | 否 | http |
| 2.27.32.81:3128 | 否 | ✓ 1702ms | ✓ 886ms | 否 | ✓ 1586ms | http |
| 222.228.171.92:8080 | ✓ 1669ms | 否 | ✓ 1627ms | ✓ 1722ms | ✓ 1196ms | http |
| 115.231.181.40:8128 | 否 | 否 | ✓ 1067ms | ✓ 1439ms | ✓ 1189ms | http |
| 59.46.216.131:30001 | ✓ 1119ms | ✓ 1621ms | ✓ 1339ms | 否 | ✓ 1354ms | http |
| 168.110.52.228:3128 | ✓ 1605ms | 否 | ✓ 1392ms | ✓ 1150ms | ✓ 1031ms | http |
| 79.132.136.58:3128 | ✓ 435ms | ✓ 1692ms | 否 | ✓ 1250ms | ✓ 919ms | http |
| 129.212.224.122:3128 | ✓ 1670ms | 否 | ✓ 927ms | ✓ 1249ms | ✓ 1000ms | http |
| 8.209.238.110:47701 | 否 | ✓ 1901ms | ✓ 1421ms | ✓ 1048ms | ✓ 1663ms | http |
| 36.103.198.235:7890 | ✓ 1978ms | 否 | ✓ 1558ms | ✓ 1519ms | 否 | http |
| 118.113.246.125:1080 | ✓ 1472ms | ✓ 1852ms | ✓ 1658ms | 否 | 否 | http |
| 137.184.0.30:3128 | ✓ 1284ms | ✓ 1297ms | ✓ 527ms | ✓ 983ms | ✓ 756ms | http |
| 95.214.9.93:3128 | ✓ 764ms | ✓ 1837ms | ✓ 650ms | ✓ 1775ms | 否 | http |
| 45.186.6.104:3128 | ✓ 1390ms | ✓ 1714ms | ✓ 1707ms | 否 | 否 | http |
| 110.42.37.202:20005 | ✓ 1402ms | ✓ 1759ms | ✓ 1447ms | ✓ 1783ms | ✓ 1756ms | http |
| 43.165.195.107:3128 | 否 | 否 | ✓ 1691ms | ✓ 1513ms | ✓ 1109ms | http |
| 150.249.255.91:3128 | ✓ 1011ms | 否 | ✓ 899ms | 否 | ✓ 1555ms | http |
| 109.107.179.140:31000 | ✓ 1653ms | ✓ 1876ms | ✓ 1183ms | 否 | 否 | http |
| 103.39.51.207:8080 | ✓ 1998ms | 否 | 否 | ✓ 1947ms | ✓ 1751ms | http |

<!-- END PROXY TABLE -->

## 🤝 参与贡献

本项目由社区驱动，欢迎任何形式的贡献。最简单的参与方式就是添加新的代理数据源。

请先阅读 **[贡献指南](CONTRIBUTING.md)** 了解如何开始。

## 🙏 支持本项目

如果您觉得本项目有帮助，欢迎给予支持，让更多人看到并参与贡献。

-   在 GitHub 上 **给本仓库加星** ⭐️
-   **分享**给朋友和同事

## ⚠️ 免责声明

-   本仓库中的代理均来自公开来源，不保证其速度、安全性或可用性。
-   使用这些代理的风险由您自行承担。
-   本仓库维护者不对任何滥用行为负责。请勿将代理用于非法用途。

## 📝 许可证

本仓库采用 MIT 许可证发布。详见 [LICENSE](LICENSE)。

## Stars
[![Star History Chart](https://api.star-history.com/svg?repos=feitianyul/free-proxy-list&type=Date)](https://star-history.com/#feitianyul/free-proxy-list&Date)
