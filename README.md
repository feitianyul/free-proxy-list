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

最后更新：2026-03-23 15:51:07 UTC（2026-03-23 23:51:07 UTC+8）

**代理总数：78**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 77 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 1 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 147.161.210.140:8800 | ✓ 1552ms | 否 | ✓ 959ms | ✓ 993ms | ✓ 879ms | http |
| 113.160.132.26:8080 | ✓ 1641ms | ✓ 1311ms | ✓ 1109ms | ✓ 1165ms | ✓ 991ms | http |
| 147.161.239.240:8800 | ✓ 1398ms | 否 | ✓ 1606ms | ✓ 1877ms | ✓ 1620ms | http |
| 167.103.34.108:8800 | ✓ 1508ms | 否 | ✓ 1622ms | ✓ 1640ms | ✓ 1368ms | http |
| 45.167.124.52:8080 | ✓ 1033ms | 否 | ✓ 1576ms | ✓ 1736ms | ✓ 1919ms | http |
| 142.171.224.229:7890 | ✓ 1693ms | ✓ 1874ms | ✓ 1129ms | ✓ 1824ms | ✓ 532ms | http |
| 120.92.212.16:8890 | ✓ 915ms | ✓ 1151ms | ✓ 977ms | ✓ 1173ms | ✓ 933ms | http |
| 121.230.9.205:1080 | ✓ 1798ms | ✓ 1574ms | ✓ 1474ms | 否 | ✓ 1105ms | http |
| 120.92.212.16:7890 | 否 | ✓ 1384ms | ✓ 1436ms | ✓ 1837ms | ✓ 1784ms | http |
| 217.76.245.80:999 | ✓ 1591ms | ✓ 1947ms | ✓ 1373ms | ✓ 1696ms | ✓ 1238ms | http |
| 35.225.22.61:80 | ✓ 781ms | 否 | ✓ 704ms | ✓ 1216ms | ✓ 957ms | http |
| 167.103.31.122:8800 | ✓ 1325ms | 否 | ✓ 1586ms | ✓ 1615ms | ✓ 1490ms | http |
| 59.46.216.131:30001 | ✓ 1864ms | ✓ 1294ms | 否 | ✓ 1532ms | 否 | http |
| 20.27.14.220:8561 | ✓ 1372ms | ✓ 1179ms | ✓ 796ms | ✓ 735ms | ✓ 656ms | http |
| 58.220.95.8:10174 | ✓ 910ms | ✓ 1582ms | 否 | 否 | ✓ 1327ms | http |
| 101.43.127.100:8877 | ✓ 1694ms | 否 | ✓ 1798ms | ✓ 1667ms | ✓ 907ms | http |
| 137.220.150.152:6005 | ✓ 1465ms | 否 | ✓ 1555ms | ✓ 1912ms | 否 | http |
| 20.78.118.91:8561 | ✓ 1340ms | ✓ 880ms | ✓ 649ms | ✓ 730ms | ✓ 653ms | http |
| 20.210.39.153:8561 | ✓ 1340ms | ✓ 1345ms | ✓ 434ms | ✓ 749ms | ✓ 575ms | http |
| 20.27.15.111:8561 | ✓ 1340ms | ✓ 1160ms | ✓ 605ms | ✓ 792ms | ✓ 576ms | http |
| 20.27.11.248:8561 | ✓ 1340ms | ✓ 1494ms | ✓ 432ms | ✓ 730ms | ✓ 571ms | http |
| 20.27.13.35:8561 | ✓ 433ms | ✓ 1302ms | ✓ 440ms | ✓ 731ms | ✓ 572ms | http |
| 20.78.26.206:8561 | ✓ 432ms | ✓ 1229ms | ✓ 441ms | ✓ 774ms | ✓ 571ms | http |
| 38.34.179.48:8445 | 否 | ✓ 1967ms | ✓ 212ms | ✓ 716ms | ✓ 848ms | http |
| 38.145.208.167:8445 | 否 | ✓ 1016ms | ✓ 682ms | 否 | ✓ 1388ms | http |
| 116.80.63.64:7777 | ✓ 1461ms | 否 | ✓ 1470ms | ✓ 1884ms | 否 | http |
| 47.101.159.19:8899 | ✓ 876ms | ✓ 1040ms | ✓ 907ms | ✓ 1051ms | ✓ 841ms | http |
| 219.117.204.211:7799 | 否 | 否 | ✓ 1941ms | ✓ 975ms | ✓ 744ms | http |
| 45.149.92.147:5001 | ✓ 594ms | 否 | ✓ 644ms | ✓ 788ms | ✓ 1101ms | http |
| 101.47.73.135:3128 | ✓ 1524ms | 否 | ✓ 1654ms | 否 | ✓ 999ms | http |
| 103.82.23.118:5207 | ✓ 1429ms | 否 | 否 | ✓ 1685ms | ✓ 1559ms | http |
| 106.75.15.167:7890 | ✓ 1150ms | 否 | 否 | ✓ 1187ms | ✓ 870ms | http |
| 137.184.6.117:3128 | ✓ 1362ms | ✓ 879ms | ✓ 663ms | ✓ 845ms | ✓ 672ms | http |
| 167.71.196.28:8080 | ✓ 1317ms | 否 | ✓ 1165ms | ✓ 1002ms | 否 | http |
| 45.136.198.40:3128 | ✓ 870ms | 否 | ✓ 847ms | ✓ 1720ms | ✓ 1358ms | http |
| 137.220.150.22:6005 | ✓ 1729ms | 否 | ✓ 1801ms | ✓ 1115ms | 否 | http |
| 38.180.2.107:3128 | ✓ 1212ms | 否 | ✓ 1145ms | 否 | ✓ 1540ms | http |
| 45.136.131.36:8450 | 否 | ✓ 913ms | ✓ 465ms | ✓ 1688ms | 否 | http |
| 101.32.244.83:8080 | ✓ 1374ms | 否 | ✓ 904ms | ✓ 1126ms | ✓ 1192ms | http |
| 121.43.196.210:8222 | ✓ 913ms | ✓ 1042ms | ✓ 805ms | ✓ 1069ms | ✓ 852ms | http |
| 121.43.196.213:8222 | ✓ 956ms | ✓ 1038ms | ✓ 803ms | ✓ 1079ms | ✓ 862ms | http |
| 114.55.226.123:10086 | ✓ 1092ms | ✓ 1700ms | ✓ 1025ms | ✓ 1293ms | ✓ 1062ms | http |
| 114.237.77.244:1080 | ✓ 1675ms | ✓ 1651ms | ✓ 875ms | ✓ 1236ms | ✓ 912ms | http |
| 88.80.150.82:8080 | ✓ 1237ms | 否 | ✓ 1038ms | 否 | ✓ 1862ms | https |
| 38.34.183.221:8451 | ✓ 599ms | 否 | ✓ 1684ms | ✓ 756ms | ✓ 567ms | http |
| 84.247.149.172:3128 | ✓ 769ms | 否 | 否 | ✓ 1811ms | ✓ 1199ms | http |
| 181.78.44.63:999 | ✓ 940ms | 否 | ✓ 1509ms | ✓ 1766ms | ✓ 1395ms | http |
| 116.80.49.161:3172 | ✓ 1884ms | 否 | ✓ 1612ms | ✓ 1771ms | 否 | http |
| 166.88.55.83:7890 | ✓ 624ms | ✓ 1444ms | ✓ 659ms | ✓ 753ms | ✓ 602ms | http |
| 210.45.70.16:7895 | ✓ 995ms | ✓ 1586ms | ✓ 1254ms | ✓ 1295ms | ✓ 1007ms | http |
| 202.141.161.53:30001 | ✓ 1097ms | ✓ 1947ms | ✓ 1007ms | ✓ 1174ms | ✓ 1332ms | http |
| 83.219.250.8:62920 | ✓ 1238ms | 否 | ✓ 1799ms | 否 | ✓ 1736ms | http |
| 38.145.218.102:8447 | 否 | ✓ 1135ms | ✓ 1182ms | ✓ 1219ms | ✓ 569ms | http |
| 104.168.158.236:10808 | 否 | 否 | ✓ 950ms | ✓ 1439ms | ✓ 1794ms | http |
| 38.34.183.234:8450 | 否 | ✓ 768ms | ✓ 769ms | ✓ 1753ms | ✓ 610ms | http |
| 45.10.69.30:8888 | ✓ 144ms | ✓ 665ms | ✓ 78ms | ✓ 727ms | ✓ 528ms | http |
| 150.249.255.91:3128 | ✓ 1408ms | ✓ 1436ms | ✓ 641ms | 否 | 否 | http |
| 38.34.178.241:8444 | ✓ 1615ms | ✓ 928ms | ✓ 318ms | ✓ 1262ms | ✓ 1651ms | http |
| 187.216.141.46:3128 | ✓ 665ms | 否 | ✓ 1019ms | ✓ 1284ms | 否 | http |
| 137.220.150.104:6005 | 否 | 否 | ✓ 1082ms | ✓ 1134ms | ✓ 830ms | http |
| 38.145.218.134:8451 | ✓ 506ms | ✓ 982ms | ✓ 148ms | ✓ 1298ms | ✓ 1101ms | http |
| 38.145.208.223:8450 | ✓ 489ms | ✓ 688ms | ✓ 338ms | ✓ 715ms | ✓ 1134ms | http |
| 38.34.183.211:8444 | ✓ 1122ms | ✓ 1137ms | ✓ 232ms | ✓ 1728ms | ✓ 1488ms | http |
| 38.34.179.86:8452 | ✓ 503ms | ✓ 766ms | ✓ 235ms | 否 | ✓ 1114ms | http |
| 38.34.183.232:8444 | ✓ 1108ms | ✓ 1179ms | ✓ 270ms | ✓ 1353ms | ✓ 1780ms | http |
| 38.145.220.6:8446 | ✓ 485ms | 否 | ✓ 151ms | ✓ 1724ms | ✓ 1493ms | http |
| 38.34.179.60:8450 | ✓ 882ms | 否 | ✓ 1308ms | ✓ 1476ms | ✓ 543ms | http |
| 103.139.138.194:3128 | ✓ 1075ms | 否 | ✓ 1606ms | ✓ 1496ms | ✓ 1111ms | http |
| 38.145.208.210:8448 | ✓ 1720ms | 否 | 否 | ✓ 844ms | ✓ 1331ms | http |
| 38.145.220.11:8444 | ✓ 1073ms | ✓ 1116ms | ✓ 105ms | ✓ 882ms | ✓ 1797ms | http |
| 51.79.207.21:8080 | ✓ 1317ms | 否 | ✓ 1645ms | ✓ 1911ms | ✓ 1133ms | http |
| 38.34.179.14:8444 | ✓ 865ms | 否 | ✓ 213ms | ✓ 851ms | ✓ 550ms | http |
| 38.34.179.11:8453 | ✓ 866ms | ✓ 874ms | ✓ 522ms | 否 | 否 | http |
| 38.145.218.102:8451 | ✓ 828ms | ✓ 1784ms | ✓ 749ms | 否 | 否 | http |
| 218.89.134.230:3333 | 否 | ✓ 1637ms | ✓ 1811ms | 否 | ✓ 1295ms | http |
| 186.148.180.46:999 | 否 | 否 | ✓ 1626ms | ✓ 1857ms | ✓ 1759ms | http |
| 47.101.149.27:9010 | 否 | ✓ 1171ms | ✓ 1290ms | ✓ 1421ms | ✓ 1328ms | http |
| 181.41.201.85:3128 | ✓ 740ms | 否 | ✓ 785ms | 否 | ✓ 1753ms | http |

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
