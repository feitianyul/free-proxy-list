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

最后更新：2026-05-03 07:17:29 UTC（2026-05-03 15:17:29 UTC+8）

**代理总数：41**

点击您需要的协议类型获取最新列表，链接始终指向最近更新的代理文件。

| 协议 | 数量 | 下载 |
|----------|-------|----------|
| HTTP | 41 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/http.txt |
| HTTPS | 0 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/https.txt |
| 通过测试 (Passed) | 41 | https://raw.githubusercontent.com/wiki/feitianyul/free-proxy-list/lists/passed.txt |

<!-- END PROXY LIST -->

以下为**通过测试**的代理前 100 条预览（五域名中任意 3 个在 2s 内通过；表格中「否」表示该域未在 2s 内成功）。完整列表请使用上方表格中的「**通过测试 (Passed)**」下载。

<!-- BEGIN PROXY TABLE -->
| 代理地址 | eastmoney.com | sse.com.cn | finance.sina.com.cn | web.ifzq.gtimg.cn | proxy.finance.qq.com | 协议 |
|----------|---------------|------------|----------------------|-------------------|---------------------|------|
| 1.231.81.166:3128 | ✓ 1708ms | ✓ 1214ms | ✓ 1209ms | ✓ 1147ms | ✓ 924ms | http |
| 34.96.238.40:8080 | ✓ 902ms | ✓ 1062ms | 否 | ✓ 982ms | ✓ 1888ms | http |
| 113.160.132.26:8080 | ✓ 1533ms | 否 | ✓ 1822ms | ✓ 1321ms | ✓ 1059ms | http |
| 148.230.4.241:999 | ✓ 1168ms | 否 | ✓ 691ms | ✓ 1606ms | ✓ 1357ms | http |
| 45.167.124.71:999 | ✓ 1451ms | ✓ 1998ms | 否 | 否 | ✓ 1817ms | http |
| 80.92.204.47:1081 | ✓ 640ms | ✓ 1585ms | ✓ 1510ms | ✓ 1897ms | ✓ 1311ms | http |
| 152.32.132.190:7890 | ✓ 1049ms | 否 | ✓ 1789ms | ✓ 793ms | ✓ 1518ms | http |
| 47.85.51.197:1080 | 否 | 否 | ✓ 1775ms | ✓ 1423ms | ✓ 1543ms | http |
| 59.46.216.131:30001 | ✓ 734ms | ✓ 1092ms | ✓ 913ms | ✓ 1015ms | ✓ 918ms | http |
| 190.12.150.244:999 | ✓ 1142ms | ✓ 1786ms | ✓ 1035ms | ✓ 1824ms | ✓ 1493ms | http |
| 212.58.132.5:8888 | ✓ 1100ms | 否 | ✓ 1096ms | ✓ 1539ms | ✓ 1225ms | http |
| 106.10.55.212:1121 | 否 | ✓ 794ms | ✓ 966ms | ✓ 1242ms | ✓ 968ms | http |
| 206.206.126.177:2412 | ✓ 1421ms | 否 | ✓ 1213ms | ✓ 996ms | ✓ 788ms | http |
| 159.223.225.118:8888 | ✓ 1292ms | 否 | ✓ 1491ms | ✓ 1984ms | 否 | http |
| 49.156.44.114:8080 | ✓ 1595ms | 否 | 否 | ✓ 1358ms | ✓ 1388ms | http |
| 109.120.156.122:8090 | ✓ 1261ms | 否 | ✓ 1090ms | 否 | ✓ 1814ms | http |
| 152.70.91.193:40000 | ✓ 1242ms | 否 | ✓ 1770ms | ✓ 1692ms | ✓ 1622ms | http |
| 101.51.166.252:8080 | ✓ 1913ms | ✓ 1970ms | ✓ 1801ms | ✓ 1440ms | ✓ 1401ms | http |
| 46.105.190.38:3128 | ✓ 1132ms | ✓ 1702ms | ✓ 652ms | 否 | ✓ 1669ms | http |
| 193.123.250.39:1080 | ✓ 710ms | 否 | ✓ 1574ms | ✓ 1501ms | ✓ 820ms | http |
| 47.77.216.82:1080 | ✓ 1124ms | ✓ 1008ms | ✓ 715ms | 否 | ✓ 618ms | http |
| 218.108.131.186:17890 | ✓ 607ms | ✓ 785ms | ✓ 686ms | ✓ 812ms | ✓ 744ms | http |
| 101.32.243.189:80 | ✓ 1215ms | 否 | ✓ 1821ms | ✓ 1124ms | 否 | http |
| 103.157.117.116:8080 | 否 | 否 | ✓ 1750ms | ✓ 1836ms | ✓ 1966ms | http |
| 121.230.8.22:1080 | 否 | ✓ 1499ms | ✓ 939ms | ✓ 1207ms | ✓ 1610ms | http |
| 220.134.5.4:8080 | ✓ 1229ms | ✓ 1308ms | ✓ 1026ms | ✓ 979ms | ✓ 997ms | http |
| 8.154.21.175:3128 | ✓ 733ms | ✓ 883ms | ✓ 698ms | ✓ 898ms | ✓ 781ms | http |
| 45.125.67.37:8443 | ✓ 834ms | 否 | ✓ 1020ms | ✓ 1136ms | ✓ 1011ms | http |
| 43.133.44.89:8888 | 否 | ✓ 1073ms | 否 | ✓ 1717ms | ✓ 1076ms | http |
| 86.104.74.110:1081 | ✓ 1552ms | ✓ 1573ms | ✓ 1858ms | 否 | 否 | http |
| 3.101.133.120:80 | 否 | ✓ 1532ms | ✓ 1564ms | ✓ 1159ms | ✓ 707ms | http |
| 8.219.97.248:80 | ✓ 1450ms | 否 | ✓ 1460ms | 否 | ✓ 1697ms | http |
| 147.45.178.211:14658 | ✓ 1424ms | 否 | ✓ 1068ms | 否 | ✓ 1996ms | http |
| 116.171.106.111:3443 | 否 | ✓ 1547ms | ✓ 1772ms | ✓ 1893ms | 否 | http |
| 168.110.52.228:3128 | ✓ 1682ms | ✓ 1647ms | ✓ 503ms | ✓ 789ms | ✓ 595ms | http |
| 23.224.193.43:3128 | 否 | ✓ 1704ms | ✓ 804ms | ✓ 803ms | ✓ 1220ms | http |
| 103.3.246.71:3128 | ✓ 913ms | 否 | ✓ 1277ms | ✓ 1161ms | ✓ 897ms | http |
| 103.250.128.8:8082 | ✓ 1282ms | 否 | ✓ 1227ms | ✓ 1342ms | ✓ 1348ms | http |
| 61.52.131.172:8443 | ✓ 1841ms | ✓ 1448ms | ✓ 754ms | ✓ 1051ms | ✓ 799ms | http |
| 117.236.124.166:3128 | ✓ 1562ms | 否 | ✓ 1664ms | 否 | ✓ 1587ms | http |
| 47.112.25.109:7890 | ✓ 1469ms | 否 | 否 | ✓ 929ms | ✓ 1100ms | http |

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
