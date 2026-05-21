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

最后更新：2026-05-21 13:10:52 UTC（2026-05-21 21:10:52 UTC+8）

**代理总数：78**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 78 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 192.99.8.15:8850 | 否 | 否 | ✓ 1111ms | ✓ 1564ms | ✓ 1489ms | http |
| 1.231.81.166:3128 | ✓ 1070ms | ✓ 1505ms | 否 | ✓ 1236ms | ✓ 1146ms | http |
| 176.111.37.216:39811 | ✓ 1122ms | ✓ 1628ms | ✓ 1810ms | 否 | ✓ 1635ms | http |
| 176.111.37.5:39811 | ✓ 942ms | ✓ 1541ms | 否 | 否 | ✓ 1732ms | http |
| 185.200.188.234:10001 | ✓ 1316ms | 否 | ✓ 1290ms | 否 | ✓ 1667ms | http |
| 8.212.167.186:8080 | ✓ 1767ms | 否 | ✓ 1175ms | ✓ 1612ms | ✓ 1129ms | http |
| 113.160.132.26:8080 | ✓ 1832ms | 否 | ✓ 1438ms | ✓ 1359ms | ✓ 1161ms | http |
| 65.109.190.168:8080 | ✓ 699ms | 否 | ✓ 1988ms | 否 | ✓ 1991ms | http |
| 138.2.92.70:8100 | ✓ 1472ms | 否 | 否 | ✓ 1548ms | ✓ 1429ms | http |
| 139.59.105.64:8080 | ✓ 1445ms | 否 | ✓ 1745ms | ✓ 1466ms | 否 | http |
| 138.2.78.251:8100 | ✓ 1458ms | 否 | 否 | ✓ 1950ms | ✓ 1831ms | http |
| 45.117.163.134:3128 | ✓ 888ms | 否 | ✓ 961ms | ✓ 1276ms | ✓ 958ms | http |
| 188.253.125.38:28798 | ✓ 888ms | 否 | ✓ 1247ms | ✓ 1319ms | ✓ 945ms | http |
| 167.86.95.198:3128 | 否 | ✓ 1533ms | ✓ 1075ms | ✓ 1713ms | ✓ 1483ms | http |
| 45.8.229.228:8080 | 否 | ✓ 1686ms | ✓ 1973ms | 否 | ✓ 1662ms | http |
| 202.28.194.139:31280 | ✓ 1681ms | 否 | ✓ 1940ms | 否 | ✓ 1947ms | http |
| 84.47.150.125:1080 | ✓ 1062ms | 否 | ✓ 1740ms | 否 | ✓ 1814ms | http |
| 2.26.92.160:3128 | ✓ 503ms | ✓ 1717ms | ✓ 640ms | 否 | 否 | http |
| 152.67.191.232:6800 | ✓ 1100ms | 否 | ✓ 1022ms | ✓ 1453ms | ✓ 1241ms | http |
| 43.130.126.146:6688 | ✓ 892ms | ✓ 1718ms | ✓ 1844ms | ✓ 1196ms | 否 | http |
| 174.137.134.182:2999 | ✓ 116ms | ✓ 968ms | ✓ 175ms | ✓ 1363ms | ✓ 804ms | http |
| 74.208.192.81:3129 | ✓ 262ms | ✓ 1604ms | ✓ 938ms | ✓ 1284ms | 否 | http |
| 152.42.170.187:9090 | 否 | 否 | ✓ 1422ms | ✓ 1213ms | ✓ 981ms | http |
| 128.199.116.219:9090 | ✓ 1491ms | 否 | 否 | ✓ 1212ms | ✓ 951ms | http |
| 146.190.80.158:9090 | ✓ 1504ms | 否 | ✓ 1572ms | ✓ 1221ms | ✓ 1014ms | http |
| 159.65.5.53:8080 | ✓ 1076ms | 否 | ✓ 1465ms | 否 | ✓ 1183ms | http |
| 5.252.33.13:2025 | ✓ 1325ms | 否 | ✓ 1309ms | 否 | ✓ 1779ms | http |
| 20.164.75.153:8080 | ✓ 1741ms | 否 | ✓ 1316ms | 否 | ✓ 1909ms | http |
| 148.230.4.241:999 | ✓ 1701ms | 否 | ✓ 1071ms | ✓ 1952ms | ✓ 1646ms | http |
| 158.255.212.55:5566 | ✓ 603ms | 否 | ✓ 1878ms | ✓ 1877ms | 否 | http |
| 158.255.212.55:9480 | ✓ 606ms | 否 | ✓ 1877ms | ✓ 1882ms | 否 | http |
| 120.92.212.16:7890 | ✓ 1143ms | ✓ 1280ms | 否 | 否 | ✓ 1850ms | http |
| 170.106.136.181:31002 | ✓ 707ms | 否 | ✓ 631ms | ✓ 817ms | 否 | http |
| 34.87.80.221:30000 | ✓ 1083ms | ✓ 1676ms | ✓ 1318ms | ✓ 1166ms | ✓ 1012ms | http |
| 195.25.20.155:3128 | ✓ 1077ms | ✓ 1898ms | ✓ 971ms | ✓ 1753ms | ✓ 1238ms | http |
| 185.104.249.25:3128 | ✓ 1878ms | 否 | ✓ 784ms | 否 | ✓ 1782ms | http |
| 144.124.227.90:21074 | ✓ 805ms | 否 | 否 | ✓ 1844ms | ✓ 1568ms | http |
| 147.45.78.89:1080 | ✓ 622ms | 否 | ✓ 1916ms | ✓ 1847ms | ✓ 1967ms | http |
| 107.150.97.83:3128 | 否 | 否 | ✓ 349ms | ✓ 1579ms | ✓ 1945ms | http |
| 114.214.165.78:10810 | ✓ 1250ms | 否 | ✓ 1412ms | ✓ 1622ms | ✓ 1331ms | http |
| 34.71.229.255:3128 | ✓ 688ms | ✓ 1207ms | ✓ 870ms | ✓ 921ms | ✓ 848ms | http |
| 103.163.132.178:3128 | ✓ 1067ms | 否 | ✓ 1803ms | 否 | ✓ 1725ms | http |
| 222.107.27.7:8017 | ✓ 1549ms | ✓ 1928ms | ✓ 1215ms | ✓ 1594ms | ✓ 1987ms | http |
| 8.154.21.175:3128 | ✓ 1234ms | 否 | ✓ 1766ms | ✓ 1270ms | ✓ 1041ms | http |
| 190.12.150.244:999 | ✓ 914ms | ✓ 1616ms | ✓ 1157ms | 否 | 否 | http |
| 144.31.73.173:3128 | ✓ 1096ms | 否 | ✓ 1254ms | ✓ 1789ms | 否 | http |
| 128.199.114.189:9090 | ✓ 1504ms | 否 | ✓ 937ms | ✓ 1269ms | ✓ 958ms | http |
| 144.31.25.69:21064 | ✓ 1132ms | 否 | ✓ 714ms | ✓ 1684ms | ✓ 1257ms | http |
| 152.42.177.32:8888 | 否 | 否 | ✓ 1573ms | ✓ 1406ms | ✓ 1429ms | http |
| 159.223.41.216:9090 | ✓ 835ms | 否 | ✓ 839ms | ✓ 1237ms | ✓ 963ms | http |
| 80.150.246.98:443 | ✓ 549ms | 否 | ✓ 1491ms | ✓ 1888ms | ✓ 1303ms | http |
| 159.89.31.62:8080 | ✓ 466ms | ✓ 1718ms | ✓ 1643ms | 否 | 否 | http |
| 167.61.202.55:3128 | ✓ 1382ms | 否 | ✓ 1267ms | 否 | ✓ 1865ms | http |
| 8.212.167.186:80 | ✓ 1195ms | 否 | ✓ 1413ms | ✓ 1299ms | ✓ 1039ms | http |
| 3.101.133.120:80 | ✓ 325ms | ✓ 1436ms | ✓ 1267ms | ✓ 1808ms | ✓ 1102ms | http |
| 157.230.2.213:3128 | ✓ 217ms | 否 | ✓ 149ms | ✓ 1190ms | ✓ 748ms | http |
| 192.163.162.111:3128 | 否 | 否 | ✓ 919ms | ✓ 1563ms | ✓ 800ms | http |
| 174.138.161.174:35197 | ✓ 1461ms | 否 | ✓ 1911ms | 否 | ✓ 1803ms | http |
| 223.205.31.179:8080 | ✓ 1706ms | 否 | 否 | ✓ 1665ms | ✓ 1632ms | http |
| 152.70.91.193:40000 | ✓ 1573ms | 否 | 否 | ✓ 1882ms | ✓ 1286ms | http |
| 210.223.44.230:3128 | ✓ 1666ms | ✓ 1949ms | 否 | 否 | ✓ 1118ms | http |
| 46.30.46.133:3128 | ✓ 1058ms | 否 | 否 | ✓ 1849ms | ✓ 1918ms | http |
| 152.32.132.190:7890 | ✓ 1837ms | 否 | ✓ 1394ms | ✓ 1438ms | ✓ 1799ms | http |
| 223.16.170.103:3128 | ✓ 1230ms | 否 | ✓ 1673ms | ✓ 1215ms | ✓ 1237ms | http |
| 62.113.119.14:8080 | ✓ 1146ms | 否 | ✓ 1398ms | ✓ 1527ms | ✓ 1158ms | http |
| 34.101.184.164:3128 | ✓ 1782ms | 否 | ✓ 1362ms | ✓ 1799ms | 否 | http |
| 174.138.161.174:8092 | ✓ 1990ms | 否 | ✓ 1406ms | 否 | ✓ 1826ms | http |
| 199.38.85.122:40028 | ✓ 837ms | 否 | ✓ 740ms | ✓ 1840ms | 否 | http |
| 31.42.164.144:23255 | ✓ 820ms | 否 | ✓ 1984ms | 否 | ✓ 1469ms | http |
| 121.230.8.80:1080 | ✓ 1985ms | ✓ 1796ms | ✓ 1620ms | 否 | ✓ 1568ms | http |
| 61.52.131.172:8443 | ✓ 986ms | 否 | ✓ 1062ms | ✓ 1277ms | ✓ 1106ms | http |
| 45.129.141.143:3128 | ✓ 1208ms | ✓ 1789ms | ✓ 1860ms | 否 | ✓ 1756ms | http |
| 115.231.181.40:8128 | ✓ 1146ms | ✓ 1272ms | 否 | 否 | ✓ 1110ms | http |
| 103.172.70.173:8080 | ✓ 1427ms | 否 | ✓ 1736ms | ✓ 1611ms | ✓ 1860ms | http |
| 103.157.117.226:81 | 否 | 否 | ✓ 1484ms | ✓ 1983ms | ✓ 1653ms | http |
| 20.78.26.206:8561 | ✓ 1096ms | ✓ 1434ms | ✓ 612ms | ✓ 1075ms | 否 | http |
| 20.78.118.91:8561 | ✓ 1094ms | ✓ 1042ms | ✓ 748ms | ✓ 949ms | ✓ 939ms | http |
| 20.27.15.111:8561 | ✓ 745ms | ✓ 1280ms | ✓ 704ms | ✓ 939ms | ✓ 780ms | http |

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
