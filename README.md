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

最后更新：2026-03-01 13:47:34 UTC（2026-03-01 21:47:34 UTC+8）

**代理总数：80**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 80 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 874ms | 否 | ✓ 1116ms | ✓ 1258ms | ✓ 787ms | http |
| 168.235.110.63:3128 | ✓ 200ms | 否 | ✓ 697ms | ✓ 1025ms | ✓ 735ms | http |
| 141.11.210.35:1080 | ✓ 482ms | 否 | ✓ 1646ms | ✓ 1252ms | ✓ 756ms | http |
| 148.135.85.87:1080 | ✓ 517ms | 否 | ✓ 1484ms | ✓ 1162ms | ✓ 1019ms | http |
| 120.238.159.230:22222 | ✓ 1124ms | ✓ 1416ms | ✓ 1240ms | ✓ 1445ms | ✓ 1049ms | http |
| 14.56.107.244:3128 | ✓ 1414ms | 否 | ✓ 1863ms | 否 | ✓ 960ms | http |
| 61.72.110.94:3128 | ✓ 1106ms | 否 | ✓ 1129ms | ✓ 1701ms | ✓ 1785ms | http |
| 120.240.35.173:22222 | 否 | ✓ 1769ms | ✓ 1131ms | ✓ 1353ms | ✓ 1082ms | http |
| 120.238.159.228:22222 | 否 | ✓ 1515ms | ✓ 1070ms | ✓ 1346ms | ✓ 1236ms | http |
| 74.208.234.198:443 | 否 | ✓ 1479ms | ✓ 1494ms | ✓ 1553ms | ✓ 1408ms | http |
| 120.92.212.16:8890 | ✓ 1471ms | ✓ 1449ms | ✓ 1512ms | ✓ 1438ms | 否 | http |
| 104.238.30.45:59741 | ✓ 1559ms | 否 | ✓ 1615ms | 否 | ✓ 1903ms | http |
| 183.249.5.117:22222 | ✓ 969ms | ✓ 1200ms | ✓ 1121ms | 否 | 否 | http |
| 104.238.30.58:63744 | ✓ 1587ms | 否 | ✓ 1743ms | 否 | ✓ 1907ms | http |
| 81.70.169.194:80 | ✓ 1173ms | 否 | ✓ 1133ms | 否 | ✓ 1225ms | http |
| 101.43.255.96:80 | ✓ 1622ms | ✓ 1415ms | ✓ 1164ms | ✓ 1563ms | 否 | http |
| 2.56.178.131:443 | ✓ 1215ms | 否 | ✓ 1758ms | ✓ 1924ms | 否 | http |
| 62.113.119.14:8080 | ✓ 1313ms | 否 | ✓ 673ms | ✓ 1712ms | ✓ 1375ms | http |
| 104.238.30.39:59741 | ✓ 1617ms | 否 | ✓ 1711ms | 否 | ✓ 1871ms | http |
| 120.232.242.119:22222 | ✓ 1092ms | ✓ 1798ms | 否 | ✓ 1603ms | ✓ 1300ms | http |
| 121.128.121.184:3128 | ✓ 1474ms | ✓ 1781ms | ✓ 1352ms | 否 | ✓ 1347ms | http |
| 35.225.22.61:80 | ✓ 708ms | 否 | ✓ 370ms | ✓ 1009ms | ✓ 924ms | http |
| 45.125.67.37:8443 | 否 | 否 | ✓ 1061ms | ✓ 1486ms | ✓ 1059ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1755ms | 否 | ✓ 1138ms | ✓ 1118ms | http |
| 120.240.35.161:22222 | ✓ 1070ms | ✓ 1505ms | ✓ 1172ms | ✓ 1413ms | 否 | http |
| 104.238.30.63:63744 | ✓ 1649ms | 否 | ✓ 1647ms | 否 | ✓ 1940ms | http |
| 139.162.46.62:3128 | ✓ 957ms | 否 | ✓ 1846ms | 否 | ✓ 1601ms | http |
| 113.59.32.141:22222 | ✓ 1438ms | ✓ 1663ms | ✓ 1405ms | ✓ 1505ms | ✓ 1293ms | http |
| 104.238.30.37:59741 | ✓ 1712ms | 否 | ✓ 1648ms | 否 | ✓ 1876ms | http |
| 91.238.104.172:2024 | ✓ 1095ms | 否 | ✓ 1147ms | 否 | ✓ 1658ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1987ms | ✓ 1595ms | 否 | ✓ 1580ms | http |
| 45.140.147.82:1082 | ✓ 430ms | ✓ 1222ms | ✓ 1282ms | 否 | 否 | http |
| 103.84.95.54:7890 | ✓ 1108ms | 否 | ✓ 1002ms | 否 | ✓ 836ms | http |
| 167.160.184.231:6005 | ✓ 252ms | 否 | ✓ 856ms | ✓ 1040ms | ✓ 1619ms | http |
| 115.231.181.40:8128 | ✓ 1223ms | ✓ 1344ms | ✓ 1469ms | 否 | 否 | http |
| 104.238.30.38:59741 | ✓ 1610ms | 否 | ✓ 1898ms | 否 | ✓ 1935ms | http |
| 20.27.15.111:8561 | ✓ 1069ms | ✓ 1009ms | ✓ 686ms | ✓ 1103ms | ✓ 824ms | http |
| 20.27.11.248:8561 | ✓ 1076ms | ✓ 1229ms | ✓ 653ms | ✓ 965ms | ✓ 781ms | http |
| 20.78.26.206:8561 | ✓ 1040ms | ✓ 1211ms | ✓ 697ms | ✓ 1072ms | ✓ 822ms | http |
| 20.27.15.49:8561 | ✓ 1065ms | ✓ 1257ms | ✓ 627ms | ✓ 1097ms | ✓ 850ms | http |
| 20.78.118.91:8561 | ✓ 1048ms | ✓ 1418ms | ✓ 678ms | ✓ 1037ms | ✓ 854ms | http |
| 20.27.14.220:8561 | ✓ 1077ms | ✓ 1695ms | ✓ 618ms | ✓ 1000ms | ✓ 792ms | http |
| 20.210.76.175:8561 | ✓ 1077ms | ✓ 1185ms | ✓ 703ms | ✓ 1104ms | ✓ 944ms | http |
| 20.210.39.153:8561 | ✓ 1090ms | 否 | ✓ 727ms | ✓ 1056ms | ✓ 781ms | http |
| 20.210.76.104:8561 | ✓ 1080ms | ✓ 1925ms | ✓ 656ms | ✓ 1044ms | ✓ 875ms | http |
| 20.210.76.178:8561 | ✓ 1096ms | 否 | ✓ 703ms | ✓ 1070ms | ✓ 917ms | http |
| 45.136.198.40:3128 | ✓ 738ms | 否 | ✓ 1472ms | ✓ 1813ms | ✓ 1517ms | http |
| 120.240.35.175:22222 | ✓ 1138ms | ✓ 1421ms | ✓ 1673ms | ✓ 1334ms | ✓ 1076ms | http |
| 90.84.188.97:8000 | ✓ 1928ms | ✓ 1640ms | 否 | 否 | ✓ 1460ms | http |
| 113.59.32.145:22222 | ✓ 1566ms | ✓ 1782ms | ✓ 1445ms | ✓ 1550ms | ✓ 1317ms | http |
| 121.230.9.75:1080 | 否 | ✓ 1599ms | ✓ 1838ms | ✓ 1719ms | 否 | http |
| 95.85.252.153:21064 | ✓ 445ms | ✓ 1661ms | ✓ 1427ms | ✓ 1968ms | 否 | http |
| 165.227.5.10:8888 | ✓ 641ms | 否 | ✓ 1225ms | ✓ 1310ms | 否 | http |
| 104.238.30.50:59741 | ✓ 1529ms | 否 | ✓ 1743ms | 否 | ✓ 1904ms | http |
| 147.45.72.6:3128 | ✓ 556ms | 否 | ✓ 915ms | ✓ 1398ms | ✓ 1238ms | http |
| 77.83.203.5:443 | ✓ 720ms | 否 | 否 | ✓ 1953ms | ✓ 1435ms | http |
| 207.254.71.62:8088 | ✓ 1012ms | 否 | ✓ 1285ms | ✓ 1836ms | 否 | http |
| 45.129.141.143:3128 | ✓ 1092ms | 否 | 否 | ✓ 1887ms | ✓ 1984ms | http |
| 104.238.30.40:59741 | ✓ 1632ms | 否 | ✓ 1871ms | 否 | ✓ 1871ms | http |
| 120.240.35.160:22222 | ✓ 1157ms | 否 | ✓ 1043ms | 否 | ✓ 1022ms | http |
| 183.249.5.214:22222 | ✓ 967ms | ✓ 1083ms | ✓ 941ms | ✓ 1143ms | ✓ 909ms | http |
| 120.240.29.51:22222 | 否 | 否 | ✓ 1178ms | ✓ 1735ms | ✓ 1530ms | http |
| 117.159.239.52:22222 | ✓ 1506ms | ✓ 1228ms | ✓ 1098ms | ✓ 1299ms | ✓ 1384ms | http |
| 36.147.78.166:80 | 否 | ✓ 1967ms | 否 | ✓ 1825ms | ✓ 1593ms | http |
| 103.104.99.29:80 | 否 | 否 | ✓ 1747ms | ✓ 1810ms | ✓ 1922ms | http |
| 103.104.99.89:80 | 否 | 否 | ✓ 1973ms | ✓ 1763ms | ✓ 1854ms | http |
| 217.77.102.18:3128 | ✓ 1503ms | 否 | ✓ 1731ms | 否 | ✓ 1525ms | http |
| 212.175.29.184:8080 | ✓ 1052ms | ✓ 1918ms | ✓ 1627ms | ✓ 1970ms | 否 | http |
| 35.234.17.221:8080 | 否 | ✓ 1901ms | ✓ 1391ms | 否 | ✓ 1209ms | http |
| 120.55.163.237:10086 | ✓ 1286ms | ✓ 1327ms | ✓ 1195ms | ✓ 1332ms | ✓ 1051ms | http |
| 120.240.35.176:22222 | 否 | ✓ 1490ms | ✓ 1089ms | ✓ 1370ms | ✓ 1102ms | http |
| 36.147.78.166:443 | ✓ 1917ms | ✓ 1908ms | 否 | 否 | ✓ 1890ms | http |
| 103.39.51.190:8080 | 否 | 否 | ✓ 1893ms | ✓ 1641ms | ✓ 1938ms | http |
| 138.124.53.25:7443 | ✓ 1000ms | 否 | ✓ 1542ms | ✓ 1781ms | 否 | http |
| 183.249.5.109:22222 | ✓ 977ms | ✓ 1348ms | 否 | ✓ 1195ms | 否 | http |
| 47.74.226.8:5001 | ✓ 1688ms | 否 | ✓ 1318ms | ✓ 1549ms | 否 | http |
| 183.249.5.111:22222 | 否 | 否 | ✓ 1036ms | ✓ 1191ms | ✓ 1139ms | http |
| 5.75.201.136:1080 | ✓ 436ms | 否 | ✓ 977ms | ✓ 1648ms | 否 | http |
| 120.240.110.112:22222 | ✓ 1335ms | ✓ 1474ms | 否 | 否 | ✓ 1161ms | http |
| 47.101.149.27:9010 | ✓ 1577ms | ✓ 1560ms | ✓ 1594ms | 否 | ✓ 1313ms | http |

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
