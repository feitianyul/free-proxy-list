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

最后更新：2026-03-02 10:50:21 UTC（2026-03-02 18:50:21 UTC+8）

**代理总数：72**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 72 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 205.209.118.30:3138 | ✓ 297ms | 否 | ✓ 981ms | ✓ 1562ms | 否 | http |
| 5.75.196.26:40000 | ✓ 1670ms | 否 | ✓ 1050ms | 否 | ✓ 1759ms | http |
| 217.76.245.80:999 | ✓ 885ms | 否 | ✓ 1177ms | ✓ 1688ms | ✓ 1353ms | http |
| 121.237.181.137:8888 | ✓ 981ms | 否 | 否 | ✓ 1324ms | ✓ 961ms | http |
| 101.43.255.96:80 | ✓ 1536ms | 否 | ✓ 1405ms | ✓ 1417ms | ✓ 1207ms | http |
| 39.104.201.40:7890 | ✓ 1057ms | 否 | ✓ 1054ms | 否 | ✓ 1279ms | http |
| 36.147.78.166:80 | 否 | ✓ 1769ms | ✓ 1768ms | ✓ 1938ms | 否 | http |
| 81.70.169.194:80 | ✓ 1126ms | 否 | ✓ 1026ms | ✓ 1422ms | ✓ 1086ms | http |
| 35.234.17.221:8080 | ✓ 935ms | ✓ 1346ms | 否 | ✓ 1341ms | 否 | http |
| 45.140.147.155:1081 | ✓ 513ms | ✓ 1407ms | ✓ 1441ms | ✓ 1727ms | ✓ 1318ms | http |
| 45.88.0.114:3128 | 否 | 否 | ✓ 557ms | ✓ 1395ms | ✓ 1081ms | http |
| 45.88.0.115:3128 | ✓ 524ms | 否 | ✓ 1956ms | ✓ 1368ms | 否 | http |
| 45.88.0.99:3128 | ✓ 1595ms | 否 | ✓ 995ms | ✓ 1362ms | ✓ 1070ms | http |
| 45.88.0.113:3128 | ✓ 508ms | 否 | 否 | ✓ 1380ms | ✓ 1061ms | http |
| 45.88.0.117:3128 | ✓ 542ms | 否 | 否 | ✓ 1390ms | ✓ 1402ms | http |
| 45.88.0.111:3128 | ✓ 1924ms | 否 | ✓ 1570ms | ✓ 1332ms | ✓ 1069ms | http |
| 45.140.147.155:1082 | ✓ 476ms | ✓ 1577ms | 否 | ✓ 1872ms | ✓ 1156ms | http |
| 45.88.0.98:3128 | ✓ 887ms | ✓ 1312ms | ✓ 1453ms | 否 | 否 | http |
| 59.46.216.131:30001 | 否 | 否 | ✓ 1211ms | ✓ 1488ms | ✓ 1178ms | http |
| 91.238.104.171:2023 | ✓ 791ms | 否 | ✓ 1208ms | ✓ 1613ms | ✓ 1223ms | http |
| 121.128.121.54:3128 | 否 | ✓ 1587ms | ✓ 1213ms | ✓ 1089ms | ✓ 888ms | http |
| 45.136.198.40:3128 | ✓ 1283ms | 否 | ✓ 1762ms | 否 | ✓ 1857ms | http |
| 120.92.212.16:8890 | ✓ 1465ms | 否 | ✓ 1334ms | ✓ 1941ms | ✓ 1784ms | http |
| 14.56.107.244:3128 | 否 | 否 | ✓ 1245ms | ✓ 1687ms | ✓ 919ms | http |
| 14.56.177.44:3128 | 否 | ✓ 1296ms | ✓ 1329ms | 否 | ✓ 1517ms | http |
| 38.207.165.2:6005 | ✓ 1221ms | 否 | 否 | ✓ 1532ms | ✓ 894ms | http |
| 41.226.37.234:3128 | ✓ 1311ms | 否 | ✓ 1990ms | 否 | ✓ 1615ms | http |
| 2.56.178.131:443 | ✓ 1278ms | 否 | ✓ 1082ms | 否 | ✓ 1931ms | http |
| 103.113.70.189:1081 | 否 | ✓ 1022ms | 否 | ✓ 1371ms | ✓ 1033ms | http |
| 185.243.218.43:49153 | ✓ 671ms | 否 | ✓ 1800ms | 否 | ✓ 1781ms | http |
| 115.76.5.32:10005 | 否 | 否 | ✓ 1782ms | ✓ 1929ms | ✓ 1633ms | http |
| 103.236.64.247:8888 | ✓ 1711ms | ✓ 1999ms | 否 | 否 | ✓ 1043ms | http |
| 101.32.244.83:8080 | ✓ 1499ms | ✓ 1875ms | ✓ 1023ms | ✓ 1304ms | ✓ 1322ms | http |
| 121.43.196.213:8222 | ✓ 1012ms | ✓ 1156ms | ✓ 952ms | ✓ 1209ms | ✓ 923ms | http |
| 121.43.196.210:8222 | ✓ 1036ms | ✓ 1145ms | ✓ 939ms | ✓ 1238ms | ✓ 938ms | http |
| 114.55.226.123:10086 | ✓ 1166ms | ✓ 1529ms | ✓ 1100ms | ✓ 1543ms | ✓ 1084ms | http |
| 94.177.131.12:3128 | ✓ 1194ms | 否 | ✓ 1215ms | ✓ 898ms | ✓ 787ms | http |
| 70.22.175.232:3128 | 否 | 否 | ✓ 1176ms | ✓ 1896ms | ✓ 1216ms | http |
| 46.249.103.192:443 | ✓ 1362ms | 否 | ✓ 1468ms | ✓ 1728ms | 否 | http |
| 61.72.110.94:3128 | ✓ 852ms | 否 | ✓ 958ms | ✓ 1417ms | 否 | http |
| 125.128.12.84:3128 | ✓ 906ms | 否 | ✓ 1953ms | 否 | ✓ 1918ms | http |
| 120.92.212.16:7890 | 否 | 否 | ✓ 1997ms | ✓ 1375ms | ✓ 1591ms | http |
| 125.128.12.194:3128 | ✓ 864ms | ✓ 1680ms | 否 | 否 | ✓ 1548ms | http |
| 221.202.27.194:10809 | 否 | ✓ 1384ms | 否 | ✓ 1962ms | ✓ 1962ms | http |
| 165.227.5.10:8888 | ✓ 738ms | 否 | ✓ 897ms | ✓ 1902ms | ✓ 1647ms | http |
| 61.72.110.54:3128 | ✓ 689ms | 否 | ✓ 1429ms | ✓ 1411ms | 否 | http |
| 125.128.12.114:3128 | ✓ 1783ms | 否 | ✓ 1786ms | ✓ 1141ms | ✓ 1858ms | http |
| 74.208.234.198:443 | ✓ 788ms | ✓ 1869ms | 否 | ✓ 1881ms | ✓ 1362ms | http |
| 103.84.95.54:7890 | ✓ 980ms | 否 | 否 | ✓ 935ms | ✓ 1078ms | http |
| 103.215.36.88:10101 | 否 | ✓ 1367ms | ✓ 1709ms | ✓ 1460ms | 否 | http |
| 45.88.0.116:3128 | ✓ 533ms | 否 | 否 | ✓ 1986ms | ✓ 1621ms | http |
| 121.40.231.103:7890 | ✓ 976ms | ✓ 1155ms | ✓ 1104ms | ✓ 1838ms | ✓ 912ms | http |
| 61.72.221.194:3128 | ✓ 1961ms | 否 | 否 | ✓ 1348ms | ✓ 1405ms | http |
| 35.225.22.61:80 | ✓ 642ms | 否 | ✓ 776ms | ✓ 1039ms | ✓ 795ms | http |
| 47.77.180.205:1080 | ✓ 543ms | 否 | ✓ 518ms | ✓ 839ms | ✓ 586ms | http |
| 183.128.208.68:7890 | ✓ 1011ms | 否 | 否 | ✓ 1863ms | ✓ 1028ms | http |
| 129.212.226.87:443 | ✓ 1254ms | 否 | ✓ 872ms | ✓ 1120ms | ✓ 877ms | http |
| 115.76.5.32:10010 | 否 | 否 | ✓ 1552ms | ✓ 1920ms | ✓ 1497ms | http |
| 61.72.221.94:3128 | ✓ 1250ms | 否 | 否 | ✓ 1601ms | ✓ 1150ms | http |
| 125.128.12.144:3128 | 否 | ✓ 1592ms | ✓ 1713ms | ✓ 1658ms | 否 | http |
| 113.165.202.31:1002 | ✓ 1492ms | 否 | ✓ 1505ms | ✓ 1429ms | 否 | http |
| 113.165.202.31:1003 | ✓ 1520ms | 否 | ✓ 1483ms | ✓ 1707ms | 否 | http |
| 91.238.104.172:2024 | ✓ 1102ms | 否 | ✓ 1682ms | 否 | ✓ 1609ms | http |
| 34.101.184.164:3128 | ✓ 1816ms | 否 | ✓ 1188ms | ✓ 1408ms | ✓ 1297ms | http |
| 165.225.113.220:10590 | 否 | 否 | ✓ 825ms | ✓ 1120ms | ✓ 899ms | http |
| 36.147.78.166:443 | 否 | ✓ 1795ms | ✓ 1711ms | ✓ 1993ms | ✓ 1564ms | http |
| 138.124.53.25:7443 | ✓ 717ms | 否 | ✓ 1187ms | ✓ 1650ms | 否 | http |
| 104.129.202.127:10810 | 否 | 否 | ✓ 925ms | ✓ 943ms | ✓ 746ms | http |
| 104.129.202.127:12354 | 否 | 否 | ✓ 926ms | ✓ 942ms | ✓ 910ms | http |
| 171.234.62.116:10002 | ✓ 1999ms | 否 | 否 | ✓ 1715ms | ✓ 1542ms | http |
| 45.140.147.82:1081 | ✓ 755ms | 否 | ✓ 1105ms | ✓ 1797ms | ✓ 1200ms | http |
| 165.225.113.220:11918 | ✓ 1624ms | 否 | ✓ 1597ms | 否 | ✓ 1957ms | http |

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
